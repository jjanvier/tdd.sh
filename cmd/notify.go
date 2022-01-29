package cmd

import (
	"errors"
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
	"strconv"
)

// The notification is displayed thanks to https://github.com/gen2brain/beeep.
// The bell is played thanks to https://github.com/faiface/beep. Beep requires Oto, which itself requires libasound2-dev on Linux.
var notifyCmd = &cobra.Command{
	Use:   "notify delay message",
	Short: "Display a notification message",
	Long: `Display a notification message that comes with a light bell sound after a given delay. 
This is used when tests have been red for too long when using the "launch" command. 
`,
	Example: "tdd notify 60 \"the tests have been red for 60 seconds... it's too long :/\"",
	Args:    cobra.ExactArgs(2),
	Hidden:  true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("The \"delay\" argument expects an integer.")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		delay, _ := strconv.Atoi(args[0])
		message := args[1]
		container.DI.NotifyHandler.HandleNotify(delay, message)
	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)
}
