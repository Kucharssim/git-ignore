/*
Copyright © 2023 SIMON KUCHARSKY <kucharssim@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Kucharssim/git-ignore/pkg/gitools"
	"github.com/spf13/cobra"
)

var here bool
var commit bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-ignore",
	Short: "Add statements to .gitignore",
	Long: `Add statements to .gitignore.
You can add multiple files or folders.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding", args, "to .gitignore")
		gitools.GitIgnore(args, here, commit)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&here, "here", false, "Write to a .gitignore file in the current working directory. Otherwise writes to the .gitignore in the root of the git repository")
	rootCmd.Flags().BoolVarP(&commit, "commit", "c", false, "Automatically commit the changes to the .gitignore file")

}
