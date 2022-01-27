package handler

import (
	"bytes"
	"github.com/jjanvier/tdd/execution"
	"github.com/jjanvier/tdd/helper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestHandleTodo(t *testing.T) {
	todoFile := helper.CreateTmpFile("")
	defer helper.RemoveTmpFile(todoFile)

	handler := _createNewHandler(todoFile.Name())
	handler.HandleTodo("here is something I have to do later")

	actual, _ := ioutil.ReadFile(todoFile.Name())
	expected := `here is something I have to do later
`
	assert.Equal(t, expected, string(actual))

	handler.HandleTodo("hmmmm, something else")

	newActual, _ := ioutil.ReadFile(todoFile.Name())
	newExpected := `here is something I have to do later
hmmmm, something else
`
	assert.Equal(t, newExpected, string(newActual))
}

func TestHandleDo(t *testing.T) {
	todoFile := helper.CreateTmpFile(`I should do that
also this should be done
really important to do that`)
	defer helper.RemoveTmpFile(todoFile)

	fakeStdin := _fakeStdinWithSecondOptionSelected()
	defer func() {
		fakeStdin.Close()
		os.Remove(fakeStdin.Name())
	}()

	handler := _createNewHandler(todoFile.Name())

	result, _ := handler.HandleDo(fakeStdin)

	assert.Equal(t, true, result.IsSuccess)
	assert.Equal(t, "git commit --allow-empty -m also this should be done", result.Command)
}

func TestHandleDone(t *testing.T) {
	todoFile := helper.CreateTmpFile(`I should do that
also this should be done
really important to do that`)
	defer helper.RemoveTmpFile(todoFile)

	handler := _createNewHandler(todoFile.Name())

	handler.HandleDone()

	actual, _ := ioutil.ReadFile(todoFile.Name())
	assert.Equal(t, "", string(actual))
}

func _createNewHandler(todoPath string) TodoHandler {
	executor := new(execution.SuccessCommandExecutorMock)
	todoList := TodoList{todoPath}
	commandFactory := execution.CommandFactory{}
	executionResultFactory := execution.ExecutionResultFactory{}
	newHandler := NewHandler{executor, commandFactory, executionResultFactory}

	return TodoHandler{todoList, newHandler, executionResultFactory}
}

// to fake stdin, we can use a temporary file as explained here https://github.com/manifoldco/promptui/issues/63#issuecomment-496871005
// to select the second option, our goal is to go down and press enter
// to go down => key "j" (that's what uses manifoldco/promptui) => ascii "106"
// to press enter => key "enter" => ascii "10" on Linux (so this test won't probably work on MacOs or Windows)
func _fakeStdinWithSecondOptionSelected() *os.File {
	buffer := bytes.Buffer{}
	buffer.Write([]byte{106, 10}) // charac
	_padForPromptUi(&buffer, 2)

	fakeStdin, _ := ioutil.TempFile("", "")
	fakeStdin.Write(buffer.Bytes())
	fakeStdin.Seek(0, 0)

	return fakeStdin
}

// manifoldco/promptui uses internally a 4096 length buffer
// which means, our fake stdin must be 4096 length
// also, that character "a" (ascii "97") is not used by promptui => so that's a character we can use to pad our buffer.
// see https://github.com/manifoldco/promptui/issues/63#issuecomment-638549034
func _padForPromptUi(buf *bytes.Buffer, size int) {
	pu := make([]byte, 4096-size)
	for i := 0; i < 4096-size; i++ {
		pu[i] = 97
	}
	buf.Write(pu)
}
