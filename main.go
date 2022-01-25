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
	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, notificationsCenter}

	Tdd(alias, conf, handler)
}

func Tdd(alias string, conf Configuration, handler AliasHandlerI) (ExecutionResult, error) {
	if "new" == alias {
		// TODO: handle when there is no message
		return handler.HandleNew(os.Args[2])
	}

	if "todo" == alias {
		// TODO: handle when there is no message
		return handler.HandleTodo(os.Args[2], todoFile)
	}

	return handler.HandleTestCommand(conf, alias)
}
