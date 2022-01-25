package main

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"os"
	"testing"
)

const commandPid = 100

type successCommandExecutorMock struct {
	mock.Mock
}

func (executor successCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return nil
}

func (executor successCommandExecutorMock) Execute(cmd Command) error {
	return nil
}

func (executor successCommandExecutorMock) ExecuteBackground(cmd Command) (int, error) {
	return commandPid, nil
}

type errorCommandExecutorMock struct {
	mock.Mock
}

func (executor errorCommandExecutorMock) ExecuteWithOutput(cmd Command) error {
	return errors.New("an error occurred during the execution of the command")
}

func (executor errorCommandExecutorMock) Execute(cmd Command) error {
	return errors.New("an error occurred during the execution of the command")
}

func (executor errorCommandExecutorMock) ExecuteBackground(cmd Command) (int, error) {
	return 0, errors.New("an error occurred during the execution of the command")
}

type notificationsCenterMock struct {
	mock.Mock
}

func (center *notificationsCenterMock) NotifyWithDelay(alias string, delay int, message string) {
	center.Called()
}

func (center *notificationsCenterMock) Reset(alias string) {
	center.Called()
}

func TestHandleAliasCommandWhenTestsPass(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	executor := new(successCommandExecutorMock)
	center := new(notificationsCenterMock)
	center.On("Reset").Once()

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, center}
	result, _ := handler.HandleTestCommand(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, true, result.IsSuccess)
	center.AssertExpectations(t)
}

func TestHandleAliasCommandWhenTestsPassAndCommitsAreAmended(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{true}}
	conf.Aliases = aliases

	executor := new(successCommandExecutorMock)
	center := new(notificationsCenterMock)
	center.On("Reset").Once()

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, center}
	result, _ := handler.HandleTestCommand(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --amend --no-edit", result.Command)
	assert.Equal(t, true, result.IsSuccess)
	center.AssertExpectations(t)
}

func TestHandleAliasCommandWhenTestsDoNotPass(t *testing.T) {
	conf := Configuration{}
	aliases := make(map[string]Alias)
	aliases["foo"] = Alias{"go test -v", 120, Git{false}}
	conf.Aliases = aliases

	executor := new(errorCommandExecutorMock)
	center := new(notificationsCenterMock)
	center.On("Reset").Once()
	center.On("NotifyWithDelay").Once()

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, center}
	result, _ := handler.HandleTestCommand(conf, "foo")

	assert.Equal(t, "go test -v && git add . && git commit --reuse-message=HEAD", result.Command)
	assert.Equal(t, false, result.IsSuccess)
	center.AssertExpectations(t)
}

func TestHandleNew(t *testing.T) {
	executor := new(successCommandExecutorMock)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, NotificationsCenter{executor, "/tmp/tdd.sh-pid-test"}}
	result, _ := handler.HandleNew("here is my commit message")

	assert.Equal(t, "git commit --allow-empty -m here is my commit message", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}

func TestCreateExecutionResultSuccess(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.CreateExecutionResultSuccess([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, true, result.IsSuccess)
}

func TestCreateExecutionResultFailure(t *testing.T) {
	factory := ExecutionResultFactory{}
	result := factory.CreateExecutionResultFailure([]Command{
		{"toto", []string{"titi", "--tata"}},
		{"foo", []string{"bar", "baz"}},
	})

	assert.Equal(t, "toto titi --tata && foo bar baz", result.Command)
	assert.Equal(t, false, result.IsSuccess)
}

func TestHandleTodo(t *testing.T) {
	executor := new(successCommandExecutorMock)

	todoFile := createTmpFile("")
	defer removeTmpFile(todoFile)

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, NotificationsCenter{executor, "/tmp/tdd.sh-pid-test"}}
	handler.HandleTodo("here is something I have to do later", todoFile.Name())

	actual, _ := ioutil.ReadFile(todoFile.Name())
	expected := `here is something I have to do later
`

	assert.Equal(t, expected, string(actual))

	handler.HandleTodo("hmmmm, something else", todoFile.Name())

	newActual, _ := ioutil.ReadFile(todoFile.Name())
	newExpected := `here is something I have to do later
hmmmm, something else
`

	assert.Equal(t, newExpected, string(newActual))
}

func TestHandleDo(t *testing.T) {
	executor := new(successCommandExecutorMock)

	todoFile := createTmpFile(`I should do that
also this should be done
really important to do that`)
	defer removeTmpFile(todoFile)

	fakeStdin := _fakeStdinWithSecondOptionSelected()
	defer func() {
		fakeStdin.Close()
		os.Remove(fakeStdin.Name())
	}()

	handler := AliasHandler{executor, CommandFactory{}, ExecutionResultFactory{}, NotificationsCenter{executor, "/tmp/tdd.sh-pid-test"}}
	result, _ := handler.HandleDo(todoFile.Name(), fakeStdin)

	assert.Equal(t, true, result.IsSuccess)
	assert.Equal(t, "git commit --allow-empty -m also this should be done", result.Command)
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
