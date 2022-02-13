package cmd

import (
	"github.com/fatih/color"
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
	"os"
)

var newCmd = &cobra.Command{
	Use:   "new purpose",
	Short: "Start a new TDD session",
	Long: `Start a new TDD session. 

The purpose of this TDD session will be used to commit your changes.`,
	Example: "tdd new \"a clear message that explains what I want to achieve\"",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res, err := container.DI.NewHandler.HandleNew(args[0])

		if err != nil {
			color.Red("❌ %s", err.Error())
			os.Exit(1)
		} else if res.IsSuccess {
			color.Green("✔ new TDD session created")
		} else {
			color.Red("❌ impossible to create a new TDD session")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
