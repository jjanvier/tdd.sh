package handler

import (
	"github.com/jjanvier/tdd/execution"
	"github.com/manifoldco/promptui"
	"io"
)

type TodoHandlerI interface {
	HandleDo(stdin io.ReadCloser) (execution.ExecutionResult, error)
	HandleTodo(message string) error
	HandleDone() error
}

type TodoHandler struct {
	Todo                   TodoList
	NewHandler             NewHandlerI
	ExecutionResultFactory execution.ExecutionResultFactory
}

func (handler TodoHandler) HandleDone() error {
	err := handler.Todo.Clear()

	return err
}

func (handler TodoHandler) HandleTodo(message string) error {
	err := handler.Todo.Add([]string{message})

	return err
}

func (handler TodoHandler) HandleDo(stdin io.ReadCloser) (execution.ExecutionResult, error) {
	todoList, err := handler.Todo.GetItems()
	if err != nil {
		return handler.ExecutionResultFactory.Failure([]execution.Command{}), err
	}

	prompt := promptui.Select{
		Label: "Here is your todo list, which task do you want to tackle?",
		Items: todoList,
		Stdin: stdin,
	}

	index, selected, err := prompt.Run()

	if selected == "" {
		return handler.ExecutionResultFactory.Failure([]execution.Command{}), err
	}

	if err != nil {
		return handler.ExecutionResultFactory.Failure([]execution.Command{}), err
	}

	err = handler.Todo.Remove(index)
	if err != nil {
		return handler.ExecutionResultFactory.Failure([]execution.Command{}), err
	}

	return handler.NewHandler.HandleNew(selected)
}
