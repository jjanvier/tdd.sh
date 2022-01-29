package cmd

import (
	"github.com/jjanvier/tdd/container"
	"github.com/jjanvier/tdd/handler"
	"github.com/spf13/cobra"
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
		alias := args[0]
		conf := handler.Load(container.ConfigurationFile)
		container.DI.AliasHandler.HandleAlias(conf, alias)
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}
