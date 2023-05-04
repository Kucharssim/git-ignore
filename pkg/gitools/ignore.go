package gitools

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// Adds terms in ignore arg to a .gitignore file. Unless here is true, the function attempts to write into the main .gitignore file.
// Otherwise it will write to a .gitignore in the current working directory.
// If commit is true, it will automatically add and commit the changes to the .gitignore.
func GitIgnore(ignore []string, here bool, commit bool) {
	filePath := ".gitignore"
	if !here {
		filePath = path.Join(GitDirRoot(), filePath)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err, "\nCould not open", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for _, line := range ignore {
		addLineToFile(line, file, scanner)
	}

	if !commit {
		return
	}

	gitAdd([]string{filePath})
	gitCommit(fmt.Sprint("Add ", ignore, " to .gitignore"))
}

// Adds a line to file, but only if the line is not already present in the file
func addLineToFile(line string, file *os.File, scanner *bufio.Scanner) {
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == line {
			fmt.Println(line, "is already included")
			return
		}
	}

	_, err := file.WriteString(line + "\n")

	if err != nil {
		fmt.Println(err, "\nCould not add", line)
	}
}
