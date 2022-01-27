package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tdd",
	Short: "A simple tool to enforce the TDD practice",
	Long: `A simple tool to enforce the TDD practice. It allows to:
  - easily launch your tests
  - automatically commit when your tests are green
  - launch a notification when your tests have been red for too long
  - have a consistent way to launch your tests across all your projects
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {

		args := rootCmd.Flags().Args()
		if len(args) != 1 {
			os.Exit(1)
		}

		// if we have only one argument, we try to launch the alias
		launchCmd.Execute()
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tdd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
