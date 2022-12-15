### Bridge — отделение абстракции от реализации

### Описание

Паттерн Bridge — это паттерн с немного загадочным определением из оригинальной книги «Gang of Four». Он отделяет абстракцию от ее реализации, так что они могут меняться независимо друг от друга. Это загадочное объяснение просто означает, что вы можете отделить даже самую базовую форму функциональности: отделить объект от того, что он делает.

Целью шаблона Bridge является придание гибкости структуре, которая часто изменяется. Знание входных и выходных данных метода позволяет нам изменять код, не зная о нем слишком много, и оставляя обеим сторонам свободу для более легкого изменения.

### Пример — два Printer и два метода Print для каждого
Для нашего примера мы перейдем к абстракции консольного принтера, чтобы упростить его. У нас будет две реализации. Первый будет писать в консоль. Вторую запись мы сделаем в интерфейс io.Writer, чтобы обеспечить большую гибкость решения. У нас также будет два абстрактных объекта-пользователя реализаций — Normal, который будет использовать каждую реализацию прямым образом, и реализация Packt, которая добавит предложение `Message from Packt:` к распечатываемому сообщению.
В конце у нас будет два объекта абстракции, которые имеют две разные реализации их функциональности. Итак, фактически у нас будет 4 возможных комбинации функциональности объектов.

Требования и критерии приемлемости:
* PrinterAPI, который принимает сообщение для печати
* Реализация API, которая просто выводит сообщение на консоль
* Реализация API, которая печатает в интерфейсе io.Writer
* Абстракция Printer с методом Print для реализации в типах печати
* Normal Printer, который реализует Printer и PrinterAPI интерфейс
* Normal Printer перенаправит сообщение непосредственно в реализацию
* Принтер Packt, который реализует абстракцию Printer и интерфейс PrinterAPI
* Принтер Packt добавит сообщение `Message from Packt:` ко всем распечаткам

### Реализация
```go
package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (d *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}
```

### Тесты
```go
package bridge

import (
	"errors"
	"strings"
	"testing"
)

func TestPrintAPI(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("error trying to use the API!"+
			" implementation: Message: %s\n", err.Error())
	}
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("content received on Writer was empty")
	return
}

func TestPrintApi2(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMsg := "you need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMsg) {
			t.Errorf("Error message was not correct.\n Actual:"+
				" %s \nExpected: %s\n", err.Error(), expectedErrorMsg)
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{Writer: &testWriter}

	expectedMsg := "Hello"
	err = api2.PrintMessage(expectedMsg)
	if err != nil {
		t.Errorf("error trying to use the API2"+
			"  implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Fatalf("API2 did not write corretly on the io.Writer."+
			" \nActual: %s \nExpected: %s\n", testWriter.Msg, expectedMsg)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMsg := "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expectedMsg,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMsg,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Errorf("the expected message on the io.Writer doesn't match actual."+
			"\nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMsg)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"
	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}
	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n"+
			"Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
```