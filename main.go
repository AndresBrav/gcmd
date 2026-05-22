package main

import (
	"fmt"
	"os"

	"git-cheatsheet/commands"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Global LipGloss styles for a premium aesthetic
var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")). // Vibrant Violet
			Padding(0, 2).
			MarginBottom(1)

	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			MarginTop(1)

	cmdNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FFCC")) // Electric Cyan/Turquoise

	categoryTagStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#04B575")). // Emerald Green tag
				Padding(0, 1)

	descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E3E3E3"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Italic(true)
)

type model struct {
	list        list.Model
	selectedCmd commands.GitCommand
	copied      bool
	quitting    bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// If the user is filtering/searching the list, let the list component handle key inputs.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			// Get the selected item from the list
			selectedItem := m.list.SelectedItem()
			if selectedItem != nil {
				cmd, ok := selectedItem.(commands.GitCommand)
				if ok {
					m.selectedCmd = cmd
					// Copy to clipboard
					err := clipboard.WriteAll(cmd.Cmd)
					if err == nil {
						m.copied = true
					}
					m.quitting = true
					return m, tea.Quit
				}
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		// Resize the list. Leave room at the bottom (10 lines) for the preview card.
		m.list.SetSize(msg.Width-h, msg.Height-v-10)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		if m.copied {
			successHeaderStyle := lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(lipgloss.Color("#04B575")).
				Padding(0, 2).
				MarginBottom(1)

			cmdStyle := lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#00FFCC"))

			return fmt.Sprintf("\n%s\n\n📌 Copied command: %s\n💡 Ready to paste with Ctrl+V and run!\n\n",
				successHeaderStyle.Render("📋 COPIED TO CLIPBOARD"),
				cmdStyle.Render(m.selectedCmd.Cmd),
			)
		}
		return "\n👋 Goodbye! Hope you find this tool useful.\n\n"
	}

	// Build the detailed preview card of the selected command
	var previewCard string
	selectedItem := m.list.SelectedItem()
	if selectedItem != nil {
		cmd, ok := selectedItem.(commands.GitCommand)
		if ok {
			tag := categoryTagStyle.Render(" " + cmd.Category + " ")
			titleLine := fmt.Sprintf("%s    %s", cmdNameStyle.Render(cmd.Cmd), tag)
			descLine := descStyle.Render(cmd.LongDesc)
			instructions := helpStyle.Render("Press [Enter] to copy to clipboard • [/] Search • [q] Quit")

			previewCard = cardStyle.Render(
				fmt.Sprintf("%s\n\n%s\n\n%s", titleLine, descLine, instructions),
			)
		}
	} else {
		previewCard = cardStyle.Render("No commands available.")
	}

	// Render title, list, and preview card below
	title := titleStyle.Render("🚀 INTERACTIVE GIT CHEATSHEET 🚀")
	return docStyle.Render(fmt.Sprintf("%s\n\n%s\n%s", title, m.list.View(), previewCard))
}

func main() {
	// 1. Load commands
	gitCmds := commands.GetCommands()
	items := make([]list.Item, len(gitCmds))
	for i, v := range gitCmds {
		items[i] = v
	}

	// 2. Create default list delegate with custom styling
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("#7D56F4")).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Bold(true)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(lipgloss.Color("#9B87F5"))

	// 3. Initialize list component
	// Default initial sizes (will be updated dynamically via WindowSizeMsg)
	cmdList := list.New(items, delegate, 80, 16)
	cmdList.Title = "Select or search for a Git command:"
	cmdList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Bold(true)

	// Hide default help since we show our own customized layout at the bottom
	cmdList.SetShowHelp(false)

	initialModel := model{
		list: cmdList,
	}

	// 4. Start Bubble Tea loop
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred while running the TUI: %v\n", err)
		os.Exit(1)
	}
}
