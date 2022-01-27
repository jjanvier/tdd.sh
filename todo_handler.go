package main

import (
	"github.com/manifoldco/promptui"
	"io"
)

type TodoHandlerI interface {
	HandleDo(stdin io.ReadCloser) (ExecutionResult, error)
	HandleTodo(message string) error
	HandleDone() error
}

type TodoHandler struct {
	todo                   TodoList
	newHandler             NewHandlerI
	executionResultFactory ExecutionResultFactory
}

func (handler TodoHandler) HandleDone() error {
	err := handler.todo.Clear()

	return err
}

func (handler TodoHandler) HandleTodo(message string) error {
	err := handler.todo.Add(message)

	return err
}

func (handler TodoHandler) HandleDo(stdin io.ReadCloser) (ExecutionResult, error) {
	todoList, err := handler.todo.GetItems()
	if err != nil {
		return handler.executionResultFactory.failure([]Command{}), err
	}

	prompt := promptui.Select{
		Label: "Here is your todo list, which task do you want to tackle?",
		Items: todoList,
		Stdin: stdin,
	}

	_, selected, err := prompt.Run()

	if selected == "" {
		return handler.executionResultFactory.failure([]Command{}), err
	}

	if err != nil {
		return handler.executionResultFactory.failure([]Command{}), err
	}

	return handler.newHandler.HandleNew(selected)
}
