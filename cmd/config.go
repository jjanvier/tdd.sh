package cmd

import (
	"github.com/fatih/color"
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
	"os"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Handle the configuration file",
	Long: `Handle the configuration file. 
You can either initialize a new configuration file or validate an existing one.`,
	Run: func(cmd *cobra.Command, args []string) {
		init, _ := cmd.Flags().GetBool("init")
		if init {
			err := container.DI.ConfigHandler.HandleInit()
			if err == nil {
				color.Green("✔ configuration file created")
			} else {
				color.Red("❌ %s", err.Error())
				os.Exit(1)
			}
		}

		validate, _ := cmd.Flags().GetBool("validate")
		if validate {
			if container.DI.ConfigHandler.HandleValidate() {
				color.Green("✔ configuration file is valid")
			} else {
				color.Red("❌ configuration file is not valid")
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("init", "i", false, "initialize a new configuration file")
	configCmd.Flags().BoolP("validate", "v", false, "validate the configuration file")
}
