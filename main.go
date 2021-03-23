package main

import (
	"log"
	"os"
)

const configurationFile = ".tdd.yml"

func main() {
	alias := os.Args[1]
	conf := Load(configurationFile)
	handler := AliasHandler{CommandExecutor{}, CommandFactory{}, ExecutionResultFactory{}}

	Tdd(alias, conf, handler)
}

func Tdd(alias string, conf Configuration, handler AliasHandlerI) ExecutionResult {
	if "new" == alias {
		if len(os.Args) < 3 {
			log.Fatal("No commit message given. Aborting.")
		}

		return handler.HandleNew(os.Args[2])
	}

	testCmd, err := conf.GetCommand(alias)
	if err != nil {
		log.Fatal(err)
	}

	return handler.HandleTestCommand(testCmd)
}

func Hello(name string) string {
	return "Hello " + name
}
