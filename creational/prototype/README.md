### Prototype — избегание создания повторяющихся объектов

### Описание

Целью шаблона Prototype является наличие объекта или набора объектов, которые уже созданы во время компиляции, с возможностью клонирования их сколько угодно раз во время выполнения. Это полезно, например, в качестве шаблона по умолчанию для пользователя, который только что зарегистрировался на вашей веб-странице, или тарифного плана по умолчанию в каком-либо сервисе. Основное различие между этим шаблоном и шаблоном Builder заключается в том, что объекты клонируются для пользователя, а не создаются во время выполнения. Вы также можете создать решение, подобное кешу, сохраняя информацию с помощью Prototype.

### Пример — магазин рубашек

Мы создадим небольшой компонент воображаемого магазина рубашек, в котором будет несколько рубашек со стандартными цветами и ценами. У каждой рубашки также будет единица складского учета (SKU - Stock Keeping Unit) — система для идентификации предметов, хранящихся в определенном месте).

Чтобы добиться того, что описано в примере, мы будем использовать прототип рубашки. Каждый раз, когда нам нужна новая рубашка, мы берем этот прототип, клонируем его и работаем с ним. 
Требования и критерии приемлемости:
* Иметь объект-прототип рубашки и интерфейс для запроса разных типов рубашек (белых, черных и синих по 15.00, 16.00 и 17.00 долларов соответственно)
* Когда вы просите белую рубашку, необходимо сделать клон белой рубашки, и новый экземпляр должен отличаться от исходного.
* Артикул (SKU) созданного объекта не должен влиять на создание нового объекта.
* Метод info должен предоставить всю информацию, доступную в полях экземпляра.


### Реализация
```go
package prototype

import "fmt"

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

type ShirtsCache struct {
}

func (sh *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, fmt.Errorf("shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs%f\n", 
		s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

```

### Тесты

```go
package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("received cache was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("type assertion for shirt couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("type assertion for shirt couldn't be done successfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to Shirt 2")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())
	t.Logf("LOG: The memory positions of the shirts are different" +
		" %p != %p\n\n", &shirt1, &shirt2)
}

```
