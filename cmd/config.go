package cmd

import (
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Handle the configuration file",
	Long: `Handle the configuration file. 
You can either initialize a new configuration file or validate an existing one.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		init, _ := cmd.Flags().GetBool("init")
		if init {
			err := container.DI.ConfigHandler.HandleInit()
			if err == nil {
				println("Configuration file created!")
			}
			return err
		}

		validate, _ := cmd.Flags().GetBool("validate")
		if validate {
			if container.DI.ConfigHandler.HandleValidate() {
				println("Configuration file is valid :)")
			} else {
				println("Configuration file is NOT valid!")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("init", "i", false, "initialize a new configuration file")
	configCmd.Flags().BoolP("validate", "v", false, "validate the configuration file")
}
