package main

import (
	"fmt"
	"testing"
)

type Command interface {
	Execute()
}

type ConcreteCommand struct {
	Receiver *Receiver
}

func (cc *ConcreteCommand) Execute() {
	cc.Receiver.Action()
}

type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver is performing an action")
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(cmd Command) {
	i.command = cmd
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func TestCommandPattern(t *testing.T) {
	receiver := &Receiver{}
	concreteCommand := &ConcreteCommand{Receiver: receiver}
	invoker := &Invoker{}
	invoker.SetCommand(concreteCommand)
	invoker.ExecuteCommand()
}

func main() {
	// run the test using command : go test -v
}
