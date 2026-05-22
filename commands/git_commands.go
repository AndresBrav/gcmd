package commands

// GitCommand represents an individual Git command
type GitCommand struct {
	Cmd       string // The command itself
	ShortDesc string // Short description for the list
	LongDesc  string // Extended explanatory description
	Category  string // Category of the command
}

// Title implements list.DefaultItem to show the title in the TUI
func (g GitCommand) Title() string {
	return g.Cmd
}

// Description implements list.DefaultItem to show the short description in the TUI
func (g GitCommand) Description() string {
	return "[" + g.Category + "] " + g.ShortDesc
}

// FilterValue implements list.Item to allow fuzzy searching over the command, description, and category
func (g GitCommand) FilterValue() string {
	return g.Cmd + " " + g.ShortDesc + " " + g.Category
}

// GetCommands returns the full list of predefined commands in English
func GetCommands() []GitCommand {
	return []GitCommand{
		// Category: Setup & Creation
		{
			Cmd:       "git init",
			ShortDesc: "Initialize a local Git repository",
			LongDesc:  "Creates a new local Git repository in the current directory. Generates the hidden '.git' folder.",
			Category:  "Setup & Creation",
		},
		{
			Cmd:       "git clone <url>",
			ShortDesc: "Clone an existing repository",
			LongDesc:  "Downloads and copies an existing remote repository from the specified URL into a local directory.",
			Category:  "Setup & Creation",
		},
		{
			Cmd:       "git config --global user.name \"<name>\"",
			ShortDesc: "Configure global username",
			LongDesc:  "Sets the author name to be associated with all your commits globally on this machine.",
			Category:  "Setup & Creation",
		},
		{
			Cmd:       "git config --global user.email \"<email>\"",
			ShortDesc: "Configure global email address",
			LongDesc:  "Sets the author email address to be associated with all your commits globally on this machine.",
			Category:  "Setup & Creation",
		},

		// Category: Saving & Status
		{
			Cmd:       "git status",
			ShortDesc: "Show current status of the working tree",
			LongDesc:  "Lists modified, staged, or untracked files compared to the last commit.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git add .",
			ShortDesc: "Stage all changes in the directory",
			LongDesc:  "Adds all new, modified, and deleted files from the current directory to the staging area for the next commit.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git add <file>",
			ShortDesc: "Stage a specific file",
			LongDesc:  "Adds changes in a specific file or directory to the staging area, preparing them for the next commit.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git commit -m \"<message>\"",
			ShortDesc: "Commit staged changes",
			LongDesc:  "Saves changes from the staging area to the local repository history with a descriptive message.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git commit --amend",
			ShortDesc: "Modify the last commit",
			LongDesc:  "Allows you to edit the message of the last commit or add files you forgot to include in it.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git stash",
			ShortDesc: "Temporarily stash uncommitted changes",
			LongDesc:  "Temporarily saves your uncommitted changes (both modified and staged) in an internal stack so you can switch branches cleanly.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git stash pop",
			ShortDesc: "Restore stashed changes",
			LongDesc:  "Removes and reapplies the most recently stashed set of changes from the stash stack.",
			Category:  "Saving & Status",
		},

		// Category: Branching & Merging
		{
			Cmd:       "git branch",
			ShortDesc: "List local branches",
			LongDesc:  "Displays all local branches. The branch you are currently on is marked with an asterisk (*).",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git branch -a",
			ShortDesc: "List all branches (local and remote)",
			LongDesc:  "Displays all local branches as well as remote tracking branches that have been downloaded.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git branch <name>",
			ShortDesc: "Create a new branch",
			LongDesc:  "Creates a new local branch pointing to the current commit, but does not switch your working directory to it.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git checkout <branch>",
			ShortDesc: "Switch to a branch or commit",
			LongDesc:  "Updates the files in your working directory to match the specified branch or commit.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git checkout -b <new-branch>",
			ShortDesc: "Create a branch and switch immediately",
			LongDesc:  "Creates a new local branch and switches your working directory to it in a single quick step.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git switch <branch>",
			ShortDesc: "Switch to a branch (Modern alternative)",
			LongDesc:  "Modern, recommended command to safely switch branches in Git.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git switch -c <new-branch>",
			ShortDesc: "Create and switch branch (Modern alternative)",
			LongDesc:  "Modern, recommended command to create a new branch and immediately switch to it.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git merge <branch>",
			ShortDesc: "Merge a branch into current branch",
			LongDesc:  "Combines the commit history and changes of the specified branch into your current working branch.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git rebase <branch>",
			ShortDesc: "Reapply your commits on top of another branch",
			LongDesc:  "Takes local commits on your current branch and reapplies them one-by one on top of the specified branch.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git branch -d <branch>",
			ShortDesc: "Delete a branch safely",
			LongDesc:  "Deletes the specified local branch. Git will prevent deletion if it contains unmerged changes.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git branch -D <branch>",
			ShortDesc: "Force delete a branch",
			LongDesc:  "Deletes the specified local branch regardless of whether its changes have been fully merged.",
			Category:  "Branching & Merging",
		},

		// Category: History & Inspection
		{
			Cmd:       "git log --oneline",
			ShortDesc: "View simplified one-line history",
			LongDesc:  "Shows a condensed list of commits containing only their abbreviated SHA hash and commit message.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git log --graph --oneline --all",
			ShortDesc: "View history tree graph",
			LongDesc:  "Draws an ASCII graphical representation of all local and remote branches along with their histories.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git diff",
			ShortDesc: "Show unstaged file differences",
			LongDesc:  "Displays line-by-line code differences between your working directory and your staging area.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git diff --staged",
			ShortDesc: "Show staged file differences",
			LongDesc:  "Displays line-by-line code differences between your staging area (staged changes) and your last commit (HEAD).",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git show <commit>",
			ShortDesc: "Show details of a specific commit",
			LongDesc:  "Displays the full commit description, author details, and the patch (diff) showing the code changes.",
			Category:  "History & Inspection",
		},

		// Category: Synchronization
		{
			Cmd:       "git remote -v",
			ShortDesc: "List configured remote repositories",
			LongDesc:  "Displays the URLs of the connected remote servers (usually 'origin') used for fetching or pushing code.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git fetch",
			ShortDesc: "Download updates from remote server",
			LongDesc:  "Fetches the history and tracking details from the remote repository to your local database without modifying your current code.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git pull",
			ShortDesc: "Download and merge remote changes",
			LongDesc:  "Fetches remote commits for the current branch and merges them immediately into your local working branch.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git push origin <branch>",
			ShortDesc: "Upload local commits to remote",
			LongDesc:  "Sends your locally committed changes to the corresponding branch on the remote server 'origin'.",
			Category:  "Synchronization",
		},

		// Category: Corrections & Reversions
		{
			Cmd:       "git reset --soft HEAD~1",
			ShortDesc: "Undo last commit, keeping changes staged",
			LongDesc:  "Undoes the last commit. All modified files remain in your staging area, ready to be committed again.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git reset --hard HEAD~1",
			ShortDesc: "Undo last commit, destroying all changes",
			LongDesc:  "Undoes the last commit and permanently destroys all code modifications made in that commit. Use with caution!",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git revert <commit>",
			ShortDesc: "Undo a commit by creating a new commit",
			LongDesc:  "Creates a new safe commit that does the exact opposite of the specified commit, ideal for public branches.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git cherry-pick <commit>",
			ShortDesc: "Apply a specific commit from another branch",
			LongDesc:  "Takes the indicated commit from any other branch and applies/copies it as a new commit onto your current branch.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git clean -fd",
			ShortDesc: "Remove untracked files and folders",
			LongDesc:  "Cleans your workspace by permanently deleting all files (-f) and directories (-d) that are not tracked by Git.",
			Category:  "Corrections & Reversions",
		},
	}
}
