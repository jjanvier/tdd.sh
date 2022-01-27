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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//newCmd.PersistentFlags().String("message", "", "The purpose of your new TDD session, which will be used as Git message")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
