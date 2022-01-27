package main

import (
	"os"
)

const configurationFile = ".tdd.yml"
const todoFile = ".tdd.todo"
const pidFile = "/tmp/tdd.sh-pid"

func main() {
	alias := os.Args[1]
	conf := Load(configurationFile)
	executor := CommandExecutor{}
	notificationsCenter := NotificationsCenter{executor, pidFile}
	todo := TodoList{todoFile}
	executionResultFactory := ExecutionResultFactory{}
	commandFactory := CommandFactory{}
	newHandler := NewHandler{executor, commandFactory, executionResultFactory}
	todoHandler := TodoHandler{todo, newHandler, executionResultFactory}
	aliasHandler := AliasHandler{executor, commandFactory, executionResultFactory, notificationsCenter}

	Tdd(alias, conf, aliasHandler, newHandler, todoHandler)
}

func Tdd(alias string, conf Configuration, aliasHandler AliasHandlerI, newHandler NewHandlerI, todoHandler TodoHandlerI) {
	if "new" == alias {
		// TODO: handle when there is no message
		newHandler.HandleNew(os.Args[2])
		return
	}

	if "todo" == alias {
		// TODO: handle when there is no message
		todoHandler.HandleTodo(os.Args[2])
		return
	}

	if "do" == alias {
		todoHandler.HandleDo(os.Stdin)
		return
	}

	if "done" == alias {
		todoHandler.HandleDone()
		return
	}

	aliasHandler.HandleAlias(conf, alias)
}
