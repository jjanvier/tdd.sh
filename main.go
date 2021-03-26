package main

import (
	"log"
	"os"
)

const configurationFile = ".tdd.yml"
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
		if len(os.Args) < 3 {
			log.Fatal("No commit message given. Aborting.")
		}

		return handler.HandleNew(os.Args[2])
	}

	return handler.HandleTestCommand(conf, alias)
}

func Hello(name string) string {
	return "Hello " + name
}
