package main

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/handler"
	"github.com/jjanvier/tdd/notification"
	"os"
)

const configurationFile = ".tdd.yml"
const todoFile = ".tdd.todo"
const pidFile = "/tmp/tdd.sh-pid"

func main() {
	alias := os.Args[1]
	conf := handler.Load(configurationFile)
	executor := execution.CommandExecutor{}
	notificationsCenter := notification.NotificationsCenter{Executor: executor, PidFileName: pidFile}
	todo := handler.TodoList{Path: todoFile}
	executionResultFactory := execution.ExecutionResultFactory{}
	commandFactory := execution.CommandFactory{}
	newHandler := handler.NewHandler{executor, commandFactory, executionResultFactory}
	todoHandler := handler.TodoHandler{todo, newHandler, executionResultFactory}
	aliasHandler := handler.AliasHandler{executor, commandFactory, executionResultFactory, notificationsCenter}

	Tdd(alias, conf, aliasHandler, newHandler, todoHandler)
}

func Tdd(alias string, conf handler.Configuration, aliasHandler handler.AliasHandlerI, newHandler handler.NewHandlerI, todoHandler handler.TodoHandlerI) {
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
