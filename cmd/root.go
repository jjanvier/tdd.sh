package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tdd",
	Short: "A simple tool to enforce the TDD practice",
	Long: `A simple tool to enforce the TDD practice. 

It follows principles erected by Kent Beck in "Test Driven Development: By Example" and allows to:
- reduce the cognitive load
	- by having a simple and consistent way to launch your tests across all your projects, whatever the language and technical stack
	- by automatically committing when your tests are green
- reduce the feedback loop by launching a notification when your tests have been red for too long
- stay focused in the red/green/refactor loop by using a todo list
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if shouldLaunchCommandBeExecuted() {
		addLaunchToOsArgs()
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}

// Try to find if "launchCommand" subcommand should be executed depending on os.Args
//
// When the command returned by rootCmd.Traverse() is the root command itself, that means
// no suitable subcommand has been found for the given os.Args. In that case, we'll try to launch the "launchCommand" to
// execute a test alias
func shouldLaunchCommandBeExecuted() bool {
	// "tdd" has been launched without any argument, so we want to the rootCommand to be executed
	if len(os.Args) == 1 {
		return false
	}

	// "tdd" has been launched only with "--help" or "-h" option, so we want to the rootCommand to be executed
	if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		return false
	}

	// "tdd" has been launched with "completion" argument, so we want to the rootCommand to be executed
	if len(os.Args) >= 2 && os.Args[1] == "completion" {
		return false
	}

	cmd, _, _ := rootCmd.Traverse(os.Args[1:])

	return cmd == rootCmd
}

// To be able to execute the "launchCommand" to execute a test alias,
// we simply add "launch" to the os.Args.
//
// For instance:
// 		- before this function, os.Args = ["tdd", "myalias"]
// 		- after this function, os.Args = ["tdd", "launch", "myalias"]
func addLaunchToOsArgs() {
	oldArgs := os.Args
	newArgs := []string{oldArgs[0], "launch"}
	newArgs = append(newArgs, oldArgs[1:]...)
	os.Args = newArgs
}
