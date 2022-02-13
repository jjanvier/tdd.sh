package cmd

import (
	"github.com/fatih/color"
	"github.com/jjanvier/tdd/container"
	"os"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Pick an item from the todo list",
	Long: `Pick an item from the todo list.

When you are ready, which means when your tests are green, you can pick a task from this list.
This will start a new TDD session by using this task as a commit message.
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := container.DI.TodoHandler.HandleDo(os.Stdin)

		if err == nil {
			color.Green("✔ new TDD session created with the item picked in the todo list")
		} else {
			color.Red("❌ %s", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
