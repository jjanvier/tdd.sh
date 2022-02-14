package execution

import (
	"fmt"
	"os/exec"
)

type CommandExecutorI interface {
	ExecuteWithOutput(cmd Command) error
	Execute(cmd Command) error
	ExecuteBackground(cmd Command) (int, error)
}

type CommandExecutor struct{}

type UnknownCommandError struct {
	command Command
}

func (e *UnknownCommandError) Error() string {
	return fmt.Sprintf("the command \"%s\" does not exist, please fix your configuration file", e.command.Name)
}

type CommandExecutionError struct {
	command Command
}

func (e *CommandExecutionError) Error() string {
	return fmt.Sprintf("error while executing the command \"%s\"", e.command.String())
}

// TODO: handle live output as explained here https://stackoverflow.com/questions/37091316/how-to-get-the-realtime-output-for-a-shell-command-in-golang
func (executor CommandExecutor) ExecuteWithOutput(cmd Command) error {
	_, err := exec.LookPath(cmd.Name)
	if err != nil {
		return &UnknownCommandError{cmd}
	}

	fmt.Println(cmd)
	out, err := exec.Command(cmd.Name, cmd.Arguments...).CombinedOutput()
	fmt.Printf("%s\n", out)

	if err != nil {
		return &CommandExecutionError{cmd}
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

func (executor CommandExecutor) ExecuteBackground(cmd Command) (int, error) {
	c := exec.Command(cmd.Name, cmd.Arguments...)
	err := c.Start()
	if err != nil {
		return -1, err
	}

	return c.Process.Pid, nil
}
