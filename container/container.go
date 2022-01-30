package container

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/handler"
	"github.com/jjanvier/tdd/notification"
)

const todoFile = ".tdd.todo"
const pidFile = "/tmp/tdd.sh-pid"
const ConfigurationFile = ".tdd.yml"

type Container struct {
	NewHandler    handler.NewHandlerI
	TodoHandler   handler.TodoHandlerI
	AliasHandler  handler.AliasHandlerI
	NotifyHandler handler.NotifyHandlerI
}

func buildDI() Container {
	commandFactory := execution.CommandFactory{}
	executor := execution.CommandExecutor{}
	notificationsCenter := notification.NotificationsCenter{Executor: executor, CommandFactory: commandFactory, PidFileName: pidFile}
	todo := handler.TodoList{Path: todoFile}
	executionResultFactory := execution.ExecutionResultFactory{}

	newHandler := handler.NewHandler{Executor: executor, CommandFactory: commandFactory, ExecutionResultFactory: executionResultFactory}
	todoHandler := handler.TodoHandler{Todo: todo, NewHandler: newHandler, ExecutionResultFactory: executionResultFactory}
	aliasHandler := handler.AliasHandler{Executor: executor, CommandFactory: commandFactory, ExecutionResultFactory: executionResultFactory, NotificationsCenter: notificationsCenter}
	notifyHandler := handler.NotifyHandler{}

	return Container{newHandler, todoHandler, aliasHandler, notifyHandler}
}

// DI global var, quite bad, but I don't know how to do best
var DI = buildDI()
