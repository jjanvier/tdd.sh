package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	Name      string
	Arguments []string
}

func (cmd Command) String() string {
	return strings.Join(append([]string{cmd.Name}, cmd.Arguments...), " ")
}

// TODO: handle live output as explained here https://stackoverflow.com/questions/37091316/how-to-get-the-realtime-output-for-a-shell-command-in-golang
func (cmd Command) ExecuteWithOutput() error {
	fmt.Println(cmd)
	out, err := exec.Command(cmd.Name, cmd.Arguments...).CombinedOutput()
	fmt.Printf("%s\n", out)
	if err != nil {
		return err
	}

	return nil
}
