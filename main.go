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

// Estilos globales de LipGloss para una estética premium
var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")). // Violeta vibrante
			Padding(0, 2).
			MarginBottom(1)

	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			MarginTop(1)

	cmdNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FFCC")) // Cyan/Turquesa eléctrico

	categoryTagStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#04B575")). // Verde esmeralda para el tag
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
		// Si el usuario está filtrando/buscando en la lista, permitimos que el componente de lista maneje las teclas.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			// Obtener el ítem seleccionado de la lista
			selectedItem := m.list.SelectedItem()
			if selectedItem != nil {
				cmd, ok := selectedItem.(commands.GitCommand)
				if ok {
					m.selectedCmd = cmd
					// Copiar al portapapeles
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
		// Redimensionar la lista. Dejamos espacio abajo (10 líneas) para la tarjeta de vista previa.
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

			return fmt.Sprintf("\n%s\n\n📌 Comando copiado: %s\n💡 ¡Listo para pegar con Ctrl+V y ejecutar!\n\n",
				successHeaderStyle.Render("📋 COPIADO AL PORTAPAPELES"),
				cmdStyle.Render(m.selectedCmd.Cmd),
			)
		}
		return "\n👋 ¡Hasta luego! Espero que te sirva esta herramienta.\n\n"
	}

	// Construir la tarjeta de vista previa detallada del comando seleccionado
	var previewCard string
	selectedItem := m.list.SelectedItem()
	if selectedItem != nil {
		cmd, ok := selectedItem.(commands.GitCommand)
		if ok {
			tag := categoryTagStyle.Render(" " + cmd.Category + " ")
			titleLine := fmt.Sprintf("%s    %s", cmdNameStyle.Render(cmd.Cmd), tag)
			descLine := descStyle.Render(cmd.LongDesc)
			instructions := helpStyle.Render("Presiona [Enter] para copiar al portapapeles • [/] Buscar • [q] Salir")

			previewCard = cardStyle.Render(
				fmt.Sprintf("%s\n\n%s\n\n%s", titleLine, descLine, instructions),
			)
		}
	} else {
		previewCard = cardStyle.Render("No hay comandos disponibles.")
	}

	// Renderizar título general, la lista y la tarjeta de vista previa abajo
	title := titleStyle.Render("🚀 GIT CHEATSHEET INTERACTIVO 🚀")
	return docStyle.Render(fmt.Sprintf("%s\n\n%s\n%s", title, m.list.View(), previewCard))
}

func main() {
	// 1. Cargar comandos
	gitCmds := commands.GetCommands()
	items := make([]list.Item, len(gitCmds))
	for i, v := range gitCmds {
		items[i] = v
	}

	// 2. Crear delegado de lista predeterminado con estilos personalizados
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("#7D56F4")).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Bold(true)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(lipgloss.Color("#9B87F5"))

	// 3. Inicializar el componente de lista
	// Dimensiones iniciales predeterminadas (se actualizan dinámicamente con WindowSizeMsg)
	cmdList := list.New(items, delegate, 80, 16)
	cmdList.Title = "Selecciona o busca un comando de Git:"
	cmdList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Bold(true)

	// Ocultar barra de ayuda del list predeterminado ya que tenemos la nuestra personalizada abajo
	cmdList.SetShowHelp(false)

	initialModel := model{
		list: cmdList,
	}

	// 4. Iniciar el loop de Bubble Tea
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ocurrió un error al ejecutar la TUI: %v\n", err)
		os.Exit(1)
	}
}
