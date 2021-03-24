package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type CommandExecutorI interface {
	ExecuteWithOutput(cmd Command) error
	Execute(cmd Command) error
}

type CommandExecutor struct {}

type Command struct {
	Name      string
	Arguments []string
}

type CommandFactory struct {}

func (factory CommandFactory) CreateGitAdd() Command {
	return Command{"git", []string{"add", "."}}
}

func (factory CommandFactory) CreateGitCommit() Command {
	return Command{"git", []string{"commit", "--reuse-message=HEAD"}}
}

func (factory CommandFactory) CreateGitCommitEmpty(message string) Command {
	return Command{"git", []string{"commit", "--allow-empty", "-m", message}}
}

func (factory CommandFactory) CreateGitCommitAmend() Command {
	return Command{"git", []string{"commit", "--amend", "--no-edit"}}
}

func (cmd Command) String() string {
	return strings.Join(append([]string{cmd.Name}, cmd.Arguments...), " ")
}

// TODO: handle live output as explained here https://stackoverflow.com/questions/37091316/how-to-get-the-realtime-output-for-a-shell-command-in-golang
func (executor CommandExecutor) ExecuteWithOutput(cmd Command) error {
	fmt.Println(cmd)
	out, err := exec.Command(cmd.Name, cmd.Arguments...).CombinedOutput()
	fmt.Printf("%s\n", out)
	if err != nil {
		return err
	}

	return nil
}

func (executor CommandExecutor) Execute(cmd Command) error {
	fmt.Println(cmd)
	err := exec.Command(cmd.Name, cmd.Arguments...).Run()
	if err != nil {
		return err
	}

	return nil
}
