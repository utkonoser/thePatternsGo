### Command — крошечный паттерн проектирования, но очень полезный

### Описание

Паттерн проектирования Command очень похож паттерн Strategy, но с ключевыми отличиями. В то время как в Strategy мы фокусируемся на изменении алгоритмов, в Command мы фокусируемся на вызове чего-либо или на абстракции некоторого типа.
Command обычно рассматривается как контейнер. Вы помещаете что-то вроде информации для взаимодействия с пользователем в пользовательский интерфейс.
При использовании Command мы пытаемся инкапсулировать какое-то действие или информацию, которые должны быть обработаны где-то еще. Это похоже на паттерн Strategy, но на самом деле Command может инициировать предварительно сконфигурированную Strategy в другом месте, так что это не одно и то же. Ниже приведены цели этого шаблона проектирования:
* Можно поместить некоторую информацию в коробку, а там, где необходимо, получатель откроет коробку и узнает ее содержимое
* Можно делегировать некоторые действия кому-то другому

### Пример — simple queue
Необходимо поместить некоторую информацию в реализацию интерфейса Command, в итоге получится очередь. Мы создадим множество экземпляров типа, реализующего паттерн Command, и передадим их в очередь, в которой будут храниться команды до тех пор, пока в очереди не окажется три из них, после чего они будут обработаны.

 Критерии приемлемости:
* Нам нужен конструктор консольных команд печати. При использовании этого конструктора со строкой он вернет команду, которая ее напечатает
* Нам нужна структура данных, которая хранит входящие команды в очереди и печатает их, когда длина очереди достигает трех
### Реализация

```go
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
```

### Еще один пример
В предыдущем примере показано, как использовать обработчик команд, который выполняет содержимое команды. Но распространенный способ использования паттерна Command — это делегирование информации вместо выполнения другому объекту.

### Реализация

```go
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
```

### Тесты

```go
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

```

