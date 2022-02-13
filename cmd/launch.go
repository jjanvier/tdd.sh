package cmd

import (
	"github.com/fatih/color"
	"github.com/jjanvier/tdd/container"
	"github.com/jjanvier/tdd/handler"
	"github.com/spf13/cobra"
	"os"
)

var launchCmd = &cobra.Command{
	Use:   "launch alias",
	Short: "Launch a test alias",
	Long: `Launch a test alias.

The alias must be present in the configuration file, in the "aliases" section. 
You can also launch this alias simply with "tdd alias".
`,
	Example: "tdd launch unit-tests",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := handler.LoadConfiguration(container.ConfigurationFile)
		if err != nil {
			color.Red("❌ %s", err.Error())
			os.Exit(1)
		}

		alias := args[0]
		res, err := container.DI.AliasHandler.HandleAlias(conf, alias)
		if err != nil {
			color.Red("❌ %s", err.Error())
			os.Exit(1)
		} else if res.IsSuccess {
			color.Green("✔ tests pass")
		} else {
			color.Red("❌ tests failed")
			os.Exit(2)
		}
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}
