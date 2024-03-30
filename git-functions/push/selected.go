package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/choose"
	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/input"
	"github.com/KevinYouu/fastGit/functions/multipleChoice"
	"github.com/KevinYouu/fastGit/git-functions/status"
)

func PushSelected() {
	fileStatuss, err := status.GetFileStatuses()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get file statuses:"), err)
		os.Exit(1)
	}

	var selectedFiles []string
	for _, fileStatus := range fileStatuss {
		if fileStatus.Status != "" {
			selectedFiles = append(selectedFiles, fileStatus.Path)
		}
	}

	data := multipleChoice.MultipleChoice(selectedFiles)
	if len(data) == 0 {
		fmt.Println(colors.RenderColor("red", "No files selected."))
		os.Exit(0)
	}

	log, err := command.RunCommand("git", "pull")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to pull: "+err.Error()))
		return
	} else {
		fmt.Println(log, colors.RenderColor("green", "Pulled successfully.\n"))
	}

	suffix := choose.Choose([]string{"fix", "feat", "docs", "style", "refactor", "test", "chore", "revert"})
	commitMessage := input.Input("Enter your commit message: \n", "commit message", "\n(esc to quit)")

	addLog, err := command.RunCommand("git", append([]string{"add"}, data...)...)
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to add files: "+err.Error()))
		return
	}
	fmt.Println(addLog, colors.RenderColor("green", "Files added successfully.\n"))

	commLog, err := command.RunCommand("git", "commit", "-m", suffix+" "+commitMessage)
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to commit: "+err.Error()))
		return
	}
	fmt.Println(commLog, colors.RenderColor("green", "Commit successful.\n"))

	pushLog, err := command.RunCommand("git", "push")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to push: "+err.Error()))
		return
	}
	fmt.Println(pushLog, colors.RenderColor("green", "Push successful."))
}