### Adapter — помощь в поддержке open/closed принципа в приложении

### Описание
Adapter очень полезен, когда, например, интерфейс устаревает и его
невозможно заменить легко или быстро. Вместо этого вы создаете новый интерфейс для удовлетворения текущих потребностей вашего приложения, которое под капотом использует реализации старого интерфейса.
Адаптер также помогает нам поддерживать принцип open/closed в наших приложениях, делая их более предсказуемыми.
Принцип open/closed впервые был сформулирован Бертраном Мейером в его книге «Object-Oriented Software Construction». Он заявил, что код должен быть открыт для новых функций, но закрыт для модификаций. Это подразумевает несколько вещей. С одной стороны, мы должны стараться писать расширяемый код, а не только работающий. В то же время мы должны стараться не модифицировать исходный код (ваш или чужой) насколько это возможно, потому что мы не всегда осознаем последствия этой модификации.

Шаблон проектирования Adapter помогает удовлетворить потребности двух частей кода, которые поначалу несовместимы. Это ключевой момент, который следует учитывать при принятии решения о том, подходит ли Adapter для решения вашей задачи.

### Пример — старый и новый Printer

В примере у нас будет старый интерфейс Printer и новый. Пользователи нового интерфейса хотят, чтобы им был доступен и старый интерфейс с дополнительной пометкой. Нам нужен Adapter, чтобы пользователи могли при необходимости использовать старые реализации (например, для работы с каким-то устаревшим кодом)

Требования и критерии приемлемости:
* Нужно создать Adapter, реализующий интерфейс ModernPrinter
* Новый объект Adapter должен содержать экземпляр интерфейса LegacyPrinter
* При использовании ModernPrinter он должен вызывать интерфейс LegacyPrinter под капотом, добавляя к нему текстовый префикс Adapter

### Реализация
```go
package adapter

import "fmt"

// legacy printer

type LegacyPrinter interface {
	Print(s string) string
}
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s", s)
	println(newMsg)
	return
}

// modern printer

type ModernPrinter interface {
	PrintStored() string
}

// printer adapter

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
```

### Тесты
```go
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

```