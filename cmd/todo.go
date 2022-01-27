package cmd

import (
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
)

var todoCmd = &cobra.Command{
	Use:   "todo item",
	Short: "Add an item to the todo list.",
	Long: `Add an item to the todo list.

While you're working on something, you can think about fixing or improving something else. 
To not loose the focus, it's handy to note the idea in a todo list.`,
	Example: "tdd todo \"hmmm... I should probably do that thing next\"",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container.DI.TodoHandler.HandleTodo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}
