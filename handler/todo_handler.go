package handler

import (
	"errors"
	"github.com/jjanvier/tdd/execution"
	"github.com/manifoldco/promptui"
	"io"
)

type TodoHandlerI interface {
	HandleDo(stdin io.ReadCloser) error
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

func (handler TodoHandler) HandleDo(stdin io.ReadCloser) error {
	todoList, err := handler.Todo.GetItems()
	if err != nil {
		return err
	}

	prompt := promptui.Select{
		Label: "Here is your todo list, which task do you want to tackle?",
		Items: todoList,
		Stdin: stdin,
	}

	index, selected, err := prompt.Run()

	if selected == "" {
		return err
	}

	if err != nil {
		return err
	}

	err = handler.Todo.Remove(index)
	if err != nil {
		return err
	}

	res, err := handler.NewHandler.HandleNew(selected)
	if err != nil {
		return err
	}

	if !res.IsSuccess {
		return errors.New("impossible to create a new TDD session")
	}

	return nil
}
