/*
Copyright © 2023 SIMON KUCHARSKY <kucharssim@gmail.com>
*/

package gitools

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func IsGitDir() bool {
	isGitCmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	isGit, err := isGitCmd.Output()

	if err != nil {
		return false
	}

	isGitString := string(isGit)
	isGitString = strings.TrimSpace(isGitString)

	isGitBool, err := strconv.ParseBool(isGitString)

	if err != nil {
		fmt.Println(err, "\nCould not understand the return of git rev-parse --is-inside-work-tree")
		os.Exit(1)
	}

	return isGitBool
}

func GitDir() string {
	isGit := IsGitDir()

	if !isGit {
		fmt.Println("It seems we are not inside of a git repository.")
		os.Exit(1)
	}

	gitDirCmd := exec.Command("git", "rev-parse", "--absolute-git-dir")

	gitDir, err := gitDirCmd.Output()

	if err != nil {
		fmt.Println(err, "\nCould not establish root of the current git repository.")
		os.Exit(1)
	}

	gitDirString := string(gitDir)
	gitDirString = strings.TrimSpace(gitDirString)

	return gitDirString
}

func GitDirRoot() string {
	gitDir := GitDir()
	gitDirRoot := path.Dir(gitDir)

	return gitDirRoot
}
