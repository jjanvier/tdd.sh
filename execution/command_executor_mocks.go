package execution

import (
	"errors"
	"github.com/stretchr/testify/mock"
)

const FakeCommandPid = 100

type SuccessCommandExecutorMock struct {
	mock.Mock
}

func (executor SuccessCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return nil
}

func (executor SuccessCommandExecutorMock) Execute(cmd Command) error {
	return nil
}

func (executor SuccessCommandExecutorMock) ExecuteBackground(cmd Command) (int, error) {
	return FakeCommandPid, nil
}

type ErrorCommandExecutorMock struct {
	mock.Mock
}

func (executor ErrorCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return errors.New("an error occurred during the execution of the command")
}

func (executor ErrorCommandExecutorMock) Execute(cmd Command) error {
	return errors.New("an error occurred during the execution of the command")
}

func (executor ErrorCommandExecutorMock) ExecuteBackground(cmd Command) (int, error) {
	return 0, errors.New("an error occurred during the execution of the command")
}

type UnknownCommandExecutorMock struct {
	mock.Mock
}

func (executor UnknownCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return &UnknownCommandError{Command{"doesnotexit", []string{}}}
}

func (executor UnknownCommandExecutorMock) Execute(cmd Command) error {
	return &UnknownCommandError{Command{"doesnotexit", []string{}}}
}

func (executor UnknownCommandExecutorMock) ExecuteBackground(cmd Command) (int, error) {
	return -1, &UnknownCommandError{Command{"doesnotexit", []string{}}}
}
