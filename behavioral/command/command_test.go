package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommandQueue(t *testing.T) {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	if len(queue.queue) != 0 {
		t.Errorf("wrong length %v must be 0", len(queue.queue))
	}
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}

func TestAnotherCommand(t *testing.T) {
	var timeCommand AnotherCommand
	timeCommand = &TimePassed{start: time.Now()}

	var helloCommand AnotherCommand
	helloCommand = &HelloMsg{}

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
