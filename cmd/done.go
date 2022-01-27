package cmd

import (
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Clear the todo list.",
	Long:  "Clear the todo list.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		container.DI.TodoHandler.HandleDone()
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
