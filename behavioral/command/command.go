package command

import (
	"fmt"
	"time"
)

// example simple queue

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutput{message: s}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}
		p.queue = make([]Command, 0)
	}
}

// The previous example shows how to use a Command handler that executes the content of
// the command. But a common way to use a Command pattern is to delegate the information,
// instead of the execution, to a different object. For example, instead of printing to
// the console, we will create a command that extracts information

type AnotherCommand interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type HelloMsg struct{}

func (h HelloMsg) Info() string {
	return "Hello World!"
}
