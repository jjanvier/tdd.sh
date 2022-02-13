package cmd

import (
	"github.com/fatih/color"
	"github.com/jjanvier/tdd/container"
	"github.com/spf13/cobra"
	"os"
)

var todoCmd = &cobra.Command{
	Use:   "todo item",
	Short: "Add an item to the todo list",
	Long: `Add an item to the todo list.

While you're working on something, you can think about fixing or improving something else. 
To not lose the focus, it's handy to note this idea in a todo list.

Please note that the purpose of this to-do list is not to store tasks that will take us months to complete. It's not an alternative to the useless ` + "`" + "// TODO: fix this" + "`" + ` that we can find sometimes in our codebase. Its goal is to minimize the red/green time: when the tests are failing, our only goal is to make them turn green quickly and easily. Anything that is not related to this precise goal should land in this todo list. Therefore, it should have a short lifetime; typically, one day maximum.
`,
	Example: "tdd todo \"hmmm... I should probably do that thing next\"",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := container.DI.TodoHandler.HandleTodo(args[0])

		if err == nil {
			color.Green("✔ item added to the todo list")
		} else {
			color.Red("❌ %s", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}
