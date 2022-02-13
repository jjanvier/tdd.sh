package cmd

import (
	"github.com/fatih/color"
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
	"os"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Clear the todo list",
	Long:  "Clear the todo list.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := container.DI.TodoHandler.HandleDone()

		if err == nil {
			color.Green("✔ todo list cleared")
		} else {
			color.Red("❌ %s", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
