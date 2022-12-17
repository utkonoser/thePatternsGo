### Decorator — старший брат паттерна Proxy

### Описание

Шаблон проектирования Decorator позволяет декорировать уже существующий тип дополнительными функциональными возможностями, фактически не касаясь его. Как это возможно? Здесь используется подход, похожий на матрешку, когда у вас есть маленькая кукла, которую вы можете поместить в куклу такой же формы, но большего размера, и так далее.
Decorator реализует тот же интерфейс, что и декорируемый им тип, и хранит экземпляр этого типа в своих полях данных. Таким образом, можно складывать столько декораторов, сколько угодно, просто сохраняя старый декоратор в поле нового.


Итак, когда именно можно использовать паттерн Decorator:
* Когда нужно добавить функциональность к некоторому коду, к которому у вас нет доступа, или вы не хотите изменять его, чтобы избежать негативного воздействия на код, и следовать принципу открытия/закрытия (например, устаревший код)
* Когда вы хотите, чтобы функциональность объекта создавалась или изменялась динамически, а количество функций неизвестно и может быстро расти

### Пример — пицца
В примере мы приготовим абстрактную пиццу, где будет пара ингредиентов для нашей пиццы – лук и мясо.

Критерии приемлемости шаблона Decorator — наличие общего интерфейса и основного типа, на основе которого будут строиться все слои:
* У нас должен быть основной интерфейс, который будут реализовывать все декораторы. Этот интерфейс будет называться `IngredientAdd`, и он будет иметь строковый метод `AddIngredient()`
* У нас должен быть основной тип `PizzaDecorator` (декоратор), к которому мы будем добавлять ингредиенты
* У нас должен быть ингредиент `Onion`, реализующий тот же интерфейс `IngredientAdd`, который добавит лук к возвращаемой пицце
* У нас должен быть ингредиент `Meat`, реализующий интерфейс `IngredientAdd`, который добавит мясо к возвращаемой пицце
* При вызове метода `AddIngredient` для верхнего объекта он должен возвращать полностью оформленную пиццу с текстом `Pizza with the following ingredients:  meat, onion`

### Реализация
```go
package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("an IngredientAdd is " +
			"needed in the Ingredient field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("an IngredientAdd is" +
			" needed in the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

```

### Тесты
```go
package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}

	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("when calling the add ingredient of the pizza"+
			" decorator it must return the text '%s'the expected text, not '%s'",
			pizzaResult, expectedText)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf(
			"when calling AddIngredient on the onion decorator without an IngredientAdd "+
				"on its Ingredient field must return an error, not a string with '%s'",
			onionResult,
		)
	}
	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()
	if err != nil {
		t.Errorf("when calling the add ingredient of the onion decorator it must "+
			"return a text with word 'onion', not '%s'", onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()
	if err == nil {
		t.Errorf(
			"when calling AddIngredient on the meat decorator without an IngredientAdd"+
				" on its Ingredient field must return an error, not a string with '%s'",
			meatResult,
		)
	}
	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	if err != nil {
		t.Errorf("when calling the add ingredient of the meat decorator it must return "+
			"a text with word 'meat', not '%s'", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("when asking for a pizza with onion and meat the returned string must"+
			" contain the text '%s' but '%s' didn't have it", expectedText, pizzaResult)
	}
	t.Log(pizzaResult)
}

```