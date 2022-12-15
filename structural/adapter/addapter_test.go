package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: Hello World!" {
		t.Errorf("message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Errorf("message didn't match: %s\n", returnedMsg)
	}
}
