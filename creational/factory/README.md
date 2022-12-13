### Factory Method — делегирование создания разных видов объектов

### Описание
При использовании шаблона проектирования Factory мы получаем дополнительный уровень инкапсуляции, чтобы наша программа могла расти в контролируемой среде. С помощью Factory Method мы делегируем создание семейств объектов другому пакету или объекту, чтобы абстрагироваться от знаний о пуле возможных объектов, которые мы могли бы использовать. Представьте, что вы хотите организовать свой отдых с помощью туристического агентства. Вы не занимаетесь гостиницами и путешествиями, а просто сообщаете агентству интересующее вас направление, чтобы оно предоставило вам все необходимое. Турагентство представляет собой Фабрику путешествий.

Цели шаблона проектирования Factory Method:
* Делегирование создания новых экземпляров структур в другую часть программы
* Работа на уровне интерфейса вместо конкретных реализаций
* Группировка семейств объектов для получения создателя объектов семейства

### Пример — Factory Method способов оплаты для магазина
В нашем примере мы собираемся реализовать метод платежей Factory, который
предоставить нам различные способы оплаты в магазине. В начале у нас будет два способа оплаты – наличные и кредитная карта. У нас также будет интерфейс с методом Pay, который должна реализовать каждая структура, используемая в качестве метода оплаты.

Требования и критерии приемлемости:
* Нужно иметь общий метод для каждого метода оплаты под названием Pay.
* Реализовать возможность делегировать создание способов оплаты Factory Method.
* Создать возможность добавлять в библиотеку дополнительные способы оплаты, просто добавляя их в Factory Method.

### Реализация
```go
package factory

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recodnized\n", m)
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

```

### Тесты
```go
package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("a payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("the cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("a payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("the debit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("a payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}

```