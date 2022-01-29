package cmd

import (
	"github.com/jjanvier/tdd/container"
	"os"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Pick an item from the todo list",
	Long: `Pick an item from the todo list.

When you are ready, which means when your tests are green, you can pickup a task from this list.
This will start a new TDD session by using this task as a commit message.
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		container.DI.TodoHandler.HandleDo(os.Stdin)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
