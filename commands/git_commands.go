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
			ShortDesc: "Set the global Git username",
			LongDesc:  "Sets the author name to be associated with all your commits globally on this machine.",
			Category:  "Setup & Creation",
		},
		{
			Cmd:       "git config --global user.email \"<email>\"",
			ShortDesc: "Set the global Git email",
			LongDesc:  "Sets the author email address to be associated with all your commits globally on this machine.",
			Category:  "Setup & Creation",
		},
		{
			Cmd:       "git config user.name",
			ShortDesc: "Show the current local username",
			LongDesc:  "Displays the Git username configured for the current local repository (not the global setting).",
			Category:  "Setup & Creation",
		},
		{
			Cmd:       "git config user.email",
			ShortDesc: "Show the current local email",
			LongDesc:  "Displays the Git email configured for the current local repository (not the global setting).",
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
			Cmd:       "git commit --allow-empty -m \"<message>\"",
			ShortDesc: "Create an empty commit without modifying files",
			LongDesc:  "Creates a commit with no file changes. Useful for marking milestones, triggering CI/CD pipelines, or adding checkpoints.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git stash",
			ShortDesc: "Temporarily save uncommitted changes",
			LongDesc:  "Temporarily saves your uncommitted changes (both modified and staged) in an internal stack so you can switch branches cleanly.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git stash list",
			ShortDesc: "Show all saved stashes",
			LongDesc:  "Lists all stash entries that have been saved with 'git stash', showing their index, branch, and message.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git stash pop",
			ShortDesc: "Restore the most recent stash",
			LongDesc:  "Removes and reapplies the most recently stashed set of changes from the stash stack.",
			Category:  "Saving & Status",
		},
		{
			Cmd:       "git rm --cached .env",
			ShortDesc: "Remove file from Git tracking without deleting it locally",
			LongDesc:  "Stops tracking the specified file in Git while keeping it on your local disk. Commonly used to remove accidentally committed secrets like .env files.",
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
			Cmd:       "git checkout main",
			ShortDesc: "Change to the main branch",
			LongDesc:  "Switches your working directory to the 'main' branch, updating all files to match the latest commit on that branch.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git checkout <branch>",
			ShortDesc: "Move to another branch",
			LongDesc:  "Updates the files in your working directory to match the specified branch or commit.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git checkout -b <new-branch>",
			ShortDesc: "Create and move to a new branch",
			LongDesc:  "Creates a new local branch and switches your working directory to it in a single quick step.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git checkout -b <local-branch> origin/<remote-branch>",
			ShortDesc: "Create a local branch linked to a remote branch",
			LongDesc:  "Creates a new local branch that tracks the specified remote branch, setting up push/pull integration automatically.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git switch <branch>",
			ShortDesc: "Alternative modern command to switch branches",
			LongDesc:  "Modern, recommended command to safely switch branches in Git.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git switch -c <new-branch>",
			ShortDesc: "Alternative modern command to create and switch branch",
			LongDesc:  "Modern, recommended command to create a new branch and immediately switch to it.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git merge <branch>",
			ShortDesc: "Merge branches",
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
		{
			Cmd:       "git branch -r",
			ShortDesc: "List all remote branches",
			LongDesc:  "Displays all remote-tracking branches that your local repository knows about. Use 'git fetch --all' first to update.",
			Category:  "Branching & Merging",
		},
		{
			Cmd:       "git branch -m <new-branch-name>",
			ShortDesc: "Rename the current branch",
			LongDesc:  "Renames your currently active local branch to the specified new name without affecting remote branches.",
			Category:  "Branching & Merging",
		},

		// Category: Remote Management
		{
			Cmd:       "git remote -v",
			ShortDesc: "Show remote repositories",
			LongDesc:  "Displays the URLs of the connected remote servers (usually 'origin') used for fetching or pushing code.",
			Category:  "Remote Management",
		},
		{
			Cmd:       "git remote add origin <repository-url>",
			ShortDesc: "Add a remote repository",
			LongDesc:  "Creates a new connection to a remote repository under the alias 'origin'. Required before you can push or pull from a new remote.",
			Category:  "Remote Management",
		},
		{
			Cmd:       "git remote remove origin",
			ShortDesc: "Remove remote repository connection",
			LongDesc:  "Deletes the remote connection named 'origin' from your local repository configuration.",
			Category:  "Remote Management",
		},
		{
			Cmd:       "git remote set-url origin git@github.com:youruser/yourrepo.git",
			ShortDesc: "Change remote URL to SSH",
			LongDesc:  "Updates the URL of the 'origin' remote to use SSH authentication instead of HTTPS, enabling key-based login.",
			Category:  "Remote Management",
		},
		{
			Cmd:       "git remote add v2 <repository-url>",
			ShortDesc: "Connect another remote repository",
			LongDesc:  "Adds a second remote repository under the alias 'v2', allowing you to push/pull from multiple sources.",
			Category:  "Remote Management",
		},
		{
			Cmd:       "git remote remove v2",
			ShortDesc: "Remove the secondary remote repository connection",
			LongDesc:  "Deletes the remote connection named 'v2' from your local repository configuration.",
			Category:  "Remote Management",
		},
		{
			Cmd:       "git remote prune <remote>",
			ShortDesc: "Remove obsolete remote-tracking references",
			LongDesc:  "Cleans up stale remote-tracking branches that no longer exist on the remote server. Run after branches are deleted remotely.",
			Category:  "Remote Management",
		},

		// Category: Synchronization
		{
			Cmd:       "git fetch",
			ShortDesc: "Download updates from remote server",
			LongDesc:  "Fetches the history and tracking details from the remote repository to your local database without modifying your current code.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git fetch --all",
			ShortDesc: "Fetch all remote branches and references",
			LongDesc:  "Downloads all branches, commits, and references from all configured remote repositories without merging them.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git fetch <remote>",
			ShortDesc: "Update remote references and data",
			LongDesc:  "Downloads all branches, commits, and references from the specified remote repository.",
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
			ShortDesc: "Push changes to the remote repository",
			LongDesc:  "Sends your locally committed changes to the corresponding branch on the remote server 'origin'.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git push origin --delete <branch>",
			ShortDesc: "Delete a remote branch from GitHub",
			LongDesc:  "Removes the specified branch from the remote repository. Your local branch remains unaffected.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git push origin <new-branch-name>",
			ShortDesc: "Push the renamed branch to the remote repository",
			LongDesc:  "Pushes the locally renamed branch up to the remote repository, creating a new remote branch with the new name.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git push origin --delete <old-branch-name>",
			ShortDesc: "Delete the old branch from the remote repository",
			LongDesc:  "After renaming a branch locally, use this to remove the old branch name from the remote repository.",
			Category:  "Synchronization",
		},
		{
			Cmd:       "git push --force",
			ShortDesc: "Force push changes to remote repository",
			LongDesc:  "Overwrites the remote branch history with your local branch. Use with extreme caution as it can erase others' work.",
			Category:  "Synchronization",
		},

		// Category: Tags
		{
			Cmd:       "git tag v1.0",
			ShortDesc: "Create a lightweight tag",
			LongDesc:  "Creates a simple tag pointing to the current commit. Lightweight tags are just named pointers with no extra metadata.",
			Category:  "Tags",
		},
		{
			Cmd:       "git tag -a v1.0 -m \"<message>\"",
			ShortDesc: "Create an annotated tag with a message",
			LongDesc:  "Creates a full tag object with author information, date, and a message. Recommended for marking release versions.",
			Category:  "Tags",
		},
		{
			Cmd:       "git tag",
			ShortDesc: "List all tags",
			LongDesc:  "Displays all tags in the repository in alphabetical order.",
			Category:  "Tags",
		},
		{
			Cmd:       "git push origin v1.0",
			ShortDesc: "Push a specific tag to the remote repository",
			LongDesc:  "Uploads a single specific tag to the remote repository. Tags are not pushed automatically with 'git push'.",
			Category:  "Tags",
		},
		{
			Cmd:       "git push origin --tags",
			ShortDesc: "Push all tags to the remote repository",
			LongDesc:  "Uploads all local tags that are not yet present on the remote repository in a single operation.",
			Category:  "Tags",
		},
		{
			Cmd:       "git checkout v1.0",
			ShortDesc: "Switch to a specific tag",
			LongDesc:  "Moves your working directory to the state of the repository at the specified tag. This puts Git in a 'detached HEAD' state.",
			Category:  "Tags",
		},
		{
			Cmd:       "git tag -d v1.0",
			ShortDesc: "Delete a local tag",
			LongDesc:  "Removes the specified tag from your local repository. The remote tag remains unaffected.",
			Category:  "Tags",
		},
		{
			Cmd:       "git push --delete origin v1.0",
			ShortDesc: "Delete a remote tag",
			LongDesc:  "Removes the specified tag from the remote repository. Your local tag remains unaffected.",
			Category:  "Tags",
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
			ShortDesc: "Visualize the branch history graph",
			LongDesc:  "Draws an ASCII graphical representation of all local and remote branches along with their histories.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git log --oneline <remote>/<branch>",
			ShortDesc: "View commits from a remote branch",
			LongDesc:  "Shows the commit history of a specific remote-tracking branch in a condensed one-line format.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git log --author=\"<name>\"",
			ShortDesc: "Search commits by author",
			LongDesc:  "Filters and displays only the commits made by the specified author name.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git log --since=\"2 weeks ago\"",
			ShortDesc: "Show commits from a specific time range",
			LongDesc:  "Displays commits made after the specified date or relative time period.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git log --grep=\"<keyword>\"",
			ShortDesc: "Search commits by message content",
			LongDesc:  "Filters and displays commits whose commit message contains the specified keyword or pattern.",
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
			ShortDesc: "Show differences between staged files and last commit",
			LongDesc:  "Displays line-by-line code differences between your staging area (staged changes) and your last commit (HEAD).",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git diff <commit1> <commit2>",
			ShortDesc: "Show differences between two commits",
			LongDesc:  "Compares the full diff of all file changes between two specified commit hashes.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git diff --name-only --diff-filter=M <commit1> <commit2>",
			ShortDesc: "Show modified files between commits",
			LongDesc:  "Lists only the names of files that were modified (not added or deleted) between two commits.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git diff --name-only --diff-filter=A <commit1> <commit2>",
			ShortDesc: "Show added files between commits",
			LongDesc:  "Lists only the names of files that were added (not modified or deleted) between two commits.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git diff --name-status <commit1> <commit2>",
			ShortDesc: "Show changed files with status between commits",
			LongDesc:  "Lists all changed files between two commits with a status letter (A=Added, M=Modified, D=Deleted).",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git diff --stat <commit1> <commit2>",
			ShortDesc: "Show statistics of changes between commits",
			LongDesc:  "Displays a summary with the number of insertions and deletions per file between two commits.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git show <commit>",
			ShortDesc: "Display detailed information about a commit",
			LongDesc:  "Displays the full commit description, author details, and the patch (diff) showing the code changes.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git reflog",
			ShortDesc: "Show the history of HEAD movements, including deleted commits",
			LongDesc:  "Displays a chronological log of all HEAD movements, including commits, resets, and checkouts. Useful for recovering deleted commits.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git blame <file>",
			ShortDesc: "Show who modified each line of a file and when",
			LongDesc:  "Annotates each line of the specified file with the commit hash, author, and date of the last modification.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git shortlog -sn",
			ShortDesc: "Display contributors and their number of commits",
			LongDesc:  "Shows a summary of commits grouped by author, sorted by number of commits. Great for seeing top contributors.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git show-branch",
			ShortDesc: "Compare branches and their unique commits",
			LongDesc:  "Shows a matrix of branches and their commits, making it easy to compare what commits are unique to each branch.",
			Category:  "History & Inspection",
		},
		{
			Cmd:       "git branch -r | grep <pattern>",
			ShortDesc: "Filter remote branches related to a pattern",
			LongDesc:  "Pipes the list of remote branches through grep to filter and display only branches matching the specified pattern.",
			Category:  "History & Inspection",
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
			ShortDesc: "Revert changes from a commit without deleting history",
			LongDesc:  "Creates a new safe commit that does the exact opposite of the specified commit, ideal for public branches.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git revert <commit> --no-edit",
			ShortDesc: "Revert a commit without opening the editor",
			LongDesc:  "Reverts the changes from a commit and auto-commits the reversion without opening the text editor for a message.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git cherry-pick <commit>",
			ShortDesc: "Apply a specific commit into the current branch",
			LongDesc:  "Takes the indicated commit from any other branch and applies/copies it as a new commit onto your current branch.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git cherry-pick <from-commit>..<to-commit>",
			ShortDesc: "Apply commits from range (from excluded, to included)",
			LongDesc:  "Applies all commits in the range from <from-commit> (excluded) to <to-commit> (included) onto the current branch.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git cherry-pick <from-commit>^..<to-commit>",
			ShortDesc: "Apply all commits including both endpoints",
			LongDesc:  "Applies all commits in the range from <from-commit> (included) to <to-commit> (included) onto the current branch.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git cherry-pick --continue",
			ShortDesc: "Continue after resolving cherry-pick conflicts",
			LongDesc:  "After manually resolving merge conflicts during a cherry-pick, use this to continue applying remaining commits.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git cherry-pick --abort",
			ShortDesc: "Cancel the cherry-pick operation",
			LongDesc:  "Aborts an in-progress cherry-pick operation and restores the branch to its state before the cherry-pick started.",
			Category:  "Corrections & Reversions",
		},
		{
			Cmd:       "git cherry-pick --skip",
			ShortDesc: "Skip the current conflicting commit",
			LongDesc:  "Skips the current conflicting commit during a cherry-pick and continues with the next commit in the range.",
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
