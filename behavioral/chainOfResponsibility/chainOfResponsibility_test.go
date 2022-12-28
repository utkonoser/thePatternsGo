package chainOfResponsibility

import (
	"fmt"
	"strings"
	"testing"
)

type myTestWriter struct {
	receivedMessage *string
}

func (m *myTestWriter) Write(p []byte) (int, error) {
	if m.receivedMessage == nil {
		m.receivedMessage = new(string)
	}
	tempMessage := fmt.Sprintf("%s%s", *m.receivedMessage, p)
	m.receivedMessage = &tempMessage
	return len(p), nil
}

func (m *myTestWriter) Next(s string) {
	m.Write([]byte(s))
}

func TestCreateDefaultChain(t *testing.T) {
	myWriter := myTestWriter{}

	writerLogger := WriterLogger{Writer: &myWriter}
	second := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds "+
		"the word 'hello', third writes to some variable if second found 'hello'",
		func(t *testing.T) {
			chain.Next("message that breaks the chain\n")

			if myWriter.receivedMessage != nil {
				t.Fatal("last link should not receive any message")
			}

			chain.Next("Hello\n")

			if myWriter.receivedMessage == nil ||
				!strings.Contains(*myWriter.receivedMessage, "Hello") {
				t.Fatal("last link didn't received expected message")
			}
		})
}
