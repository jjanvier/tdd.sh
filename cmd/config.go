package cmd

import (
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Handle the configuration file",
	Long:  `Handle the configuration file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		init, _ := cmd.Flags().GetBool("init")
		if init {
			err := container.DI.ConfigHandler.HandleInit()
			if err == nil {
				print("Configuration file created!")
			}
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("init", "i", false, "initialize a new configuration file")
}
