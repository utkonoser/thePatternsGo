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
