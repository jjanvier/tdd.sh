package main

import (
	"log"
	"os"
)

const configurationFile = ".tdd.yml"

func main() {
	alias := os.Args[1]
	conf := Load(configurationFile)
	testCmd, err := conf.GetCommand(alias)
	if err != nil {
		log.Fatal(err)
	}
	handler := AliasHandler{CommandExecutor{}, CommandFactory{}, ExecutionResultFactory{}}
	handler.HandleTestCommand(testCmd)
}

func Hello(name string) string {
	return "Hello " + name
}
