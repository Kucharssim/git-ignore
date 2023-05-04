package gitools

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func gitAdd(contents []string) {
	addCmd := exec.Command("git", "add", strings.Join(contents, " "))

	err := addCmd.Run()
	if err != nil {
		fmt.Println("Could not add files")
		os.Exit(1)
	}
}

func gitCommit(message string) {
	commitCmd := exec.Command("git", "commit", "-m", message)

	err := commitCmd.Run()

	if err != nil {
		fmt.Println("Could not commit files")
		os.Exit(1)
	}
}
