### Chain of responsibility — цепочка ответственности

### Описание

Как следует из названия, паттерн состоит из цепочки, и в нашем случае каждое звено этой цепочки подчиняется принципу единственной ответственности.
Принцип единственной ответственности подразумевает, что тип, функция, метод или любая подобная абстракция должны иметь только одну единственную ответственность. Таким образом, мы можем применять множество функций, каждая из которых выполняет одну конкретную задачу к некоторой структуре, срезу, мапе и так далее.
Один из очевидных примеров данного паттерна — это logging chain (цепочка логов).

Цепочка логов — это набор типов, которые регистрируют выходные данные некоторой программы более чем в одном интерфейсе `io.Writer`. Возможны логгеры, которые выводят логи на консоль, либо в файл, либо на удаленный сервер. Можно делать три вызова каждый раз, когда нужно записать все логи, но более элегантно сделать только один вызов и спровоцировать цепную реакцию. Но также мы могли бы иметь цепочку проверок и в случае сбоя одной из них разорвать цепочку и что-то вернуть. Это работа промежуточного программного обеспечения аутентификации и авторизации.

### Пример — multi-logger chain

Мы собираемся разработать решение с несколькими логгерами, которое мы можем связать так, как мы хотим. Мы будем использовать два разных консольных логгера и один логгер общего назначения.

Критерии приемлемости:
* Нужен простой логгер, который логирует текст запроса с префиксом `First logger` и передает его следующему звену в цепочке
* Второй логгер напишет на консоль, если во входящем тексте есть слово `hello`, и передаст запрос третьему логгеру. Но, если нет, то цепочка порвется
* Третий тип логгера общего назначения под названием `WriterLogger`, который использует для записи интерфейс `io.Writer`
* Конкретная реализация `WriterLogger` записывает в файл и представляет собой третье звено в цепочке

### Реализация

```go
package chainOfResponsibility

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First Logger: %s\n", s)
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (f *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second Logger: %s", s)

		if f.NextChain != nil {
			f.NextChain.Next(s)
		}
		return
	}
	fmt.Printf("Finishing in second logging\n\n")

}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("Writer Logger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}


```

### Тесты

```go
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

```

