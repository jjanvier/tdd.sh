package cmd

import (
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new purpose",
	Short: "Start a new TDD session",
	Long: `Start a new TDD session. 

The purpose of this TDD session will be used to commit your changes.`,
	Example: "tdd new \"a clear message that explains what I want to achieve\"",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		container.DI.NewHandler.HandleNew(args[0])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
