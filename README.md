## The Patterns - Go
Паттерны, реализованные в Go с примерами для обучения.

Репозиторий представляет собой набор реализаций с открытым исходным кодом различных паттернов, реализованных в Go.
*******************************************
### Список паттернов
*******************************************
#### Порождающие (Creational) 
<details><summary> Singleton</summary>

### Singleton — наличие уникального экземпляра типа во всей программе

### Описание

Шаблон Singleton легко запомнить. Как следует из названия, он предоставляет единственный экземпляр объекта и гарантирует отсутствие дубликатов.
При первом вызове экземпляра он создается, а затем повторно используется всеми частями в приложении, которое должно использовать это конкретное поведение.
Шаблон Singleton используется во многих различных ситуациях. Например:
* Если вы хотите использовать одно и то же соединение с базой данных для выполнения каждого запроса.
* Когда вы открываете соединение Secure Shell (SSH) с сервером для выполнения нескольких задач, и не хотите заново открывать соединение для каждой задачи.
* Если вам нужно ограничить доступ к какой-либо переменной или пространству, вы используете Singleton как дверь к этой переменной.


### Пример — уникальный счетчик
В качестве примера объекта, будет уникальный счетчик, для которого мы должны убедиться, что существует только один такой экземпляр, счетчик будет содержать количество вызовов во время исполнения программы. Неважно, сколько у нас экземпляров счетчика, все они будут считать одно и то же значение, и оно должно быть согласовано между экземплярами.

Требования и критерии приемлемости:
* Если счетчик ранее не создавался, создается новый со значением 0.
* Если счетчик уже создан, возвращается экземпляр, содержащий фактический
  счетчик.
* Если мы вызываем метод `AddOne()`, счетчик должен быть увеличен на 1.

### Реализация

```go
package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
```

### Тесты

```go
package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("expected pointer to Singleton after " +
			"calling GetInstance(), not nil\n")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("after AddOne() the count must be" +
			" 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("expected same instance in counter2 but" +
			" it got a different instance\n")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("after AddOne() the count must be" +
			" 2 but it is %d\n", currentCount)
	}
}

```


</details>

<details><summary> Builder</summary>

### Builder — повторное использование алгоритма для создания множества реализаций интерфейса.

### Описание

Шаблон Builder помогает нам создавать сложные объекты без непосредственного создания их структуры или написания необходимой логики. Представьте себе объект, который может иметь десятки полей, сами по себе являющимися более сложными структурами. Теперь представьте, что у вас есть много объектов с такими характеристиками. Здесь и пригодится Builder, чтобы не писать логику для создания всех этих объектов.

Шаблон проектирования Builder пытается:
* Абстрагировать сложные создания, чтобы создание объекта было отделено от пользователя объекта.
* Создать объект шаг за шагом, заполнив его поля и создав встроенные объекты.
* Реализовать повторное использование алгоритма создания объекта между многими объектами

### Пример – производство автомобилей
Шаблон проектирования Builder обычно описывается как отношения между Директором, несколькими Строителями и Продуктом, который они создают. Мы создадим конструктор транспортных средств. Процесс создания транспортного средства (продукта) более или менее одинаков для всех видов транспортных средств — нужно выбрать тип транспортного средства, собрать конструкцию, поместить колеса и расставить сиденья. Мы построим автомобиль и мотоцикл (два Строителя) с этим описанием. В примере директор представлен типом `ManufacturingDirector`.

Требования и критерии приемлемости:
* Должен быть производственный тип, который строит все, что нужно транспортному средству.
* При использовании сборщика автомобилей необходимо вернуть VehicleProduct с четырьмя колесами, пятью сиденьями и структурой, определенной как Car.
* При использовании сборщика мотоциклов необходимо вернуть VehicleProduct с двумя колесами, двумя сиденьями и структурой, определенной как Motorbike.
* VehicleProduct, созданный любым компоновщиком BuildProcess, должен быть открыт для модификаций.

### Реализация
```go
package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}

```
### Тесты
```go
package builder

import "testing"

func TestCarBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("wheels on a car must be 4"+
			" and they were %d \n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("structure on a car must be "+
			"'Car' and was %s \n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("seats on a car must be 5"+
			" and they were %d \n", car.Seats)
	}
}

func TestBikeBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.Build()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("wheels on a motorbike must be 2"+
			" and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must"+
			" be 'Motorbike' and was %s\n", motorbike.Structure)
	}
}

```
</details>

<details><summary> Factory Method</summary>

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
</details>

<details><summary> Abstract Factory</summary>

### Abstract Factory – фабрика фабрик

### Описание
Шаблон проектирования Abstract Factory — это новый уровень группировки для получения более крупного (и более сложного) составного объекта, который используется через его интерфейсы. Идея группировки объектов в семейства и группирования семейств состоит в том, чтобы иметь большие фабрики, которые можно было бы взаимозаменяемо и легче расширять.

Группировка связанных семейств объектов очень удобна, когда количество объектов растет настолько, что создание уникальной точки для их всех кажется единственным способом добиться гибкости создания объектов во время выполнения. Вам должны быть ясны следующие цели метода абстрактной фабрики:
* Обеспечение нового уровня инкапсуляции для фабричных методов, которые возвращают общий интерфейс для всех фабрик.
* Группировка обычных фабрик в суперфабрику (также называемую фабрикой фабрик).

### Пример – автозавод по производству мотоциклов и машин

В примере мы собираемся повторно использовать фабрику, которую создали в шаблоне проектирования Builder. В конце концов, результатом будет являться фабрика фабрик мотоциклов и машин, которые в свою очередь будут производить различные виды мотоциклов и машин соответственно.

Требования и критерии приемлемости:
* Мы должны получить объект Vehicle, используя фабрику, возвращенную абстрактной фабрикой.
* Транспортное средство должно быть конкретной реализацией мотоцикла или автомобиля, которая реализует оба интерфейса (транспортное средство и автомобиль или транспортное средство и мотоцикл).

### Реализация

```go
package abstractFactory

import "fmt"

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized \n", f)
	}
}

// Factory of factories

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized\n", v)
	}
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized\n", v)
	}
}

// for Car Factory

type Car interface {
	NumDoors() int
}

type LuxuryCar struct{}

func (*LuxuryCar) NumDoors() int {
	return 4
}

func (*LuxuryCar) NumWheels() int {
	return 4
}

func (*LuxuryCar) NumSeats() int {
	return 5
}

type FamilyCar struct{}

func (*FamilyCar) NumDoors() int {
	return 5
}

func (*FamilyCar) NumWheels() int {
	return 4
}

func (*FamilyCar) NumSeats() int {
	return 5
}

// For Motorbike Factory

type Motorbike interface {
	GetMotorbikeType() int
}

type SportMotorbike struct{}

func (s *SportMotorbike) NumWheels() int {
	return 2
}

func (s *SportMotorbike) NumSeats() int {
	return 1
}

func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) NumWheels() int {
	return 2
}

func (c *CruiseMotorbike) NumSeats() int {
	return 2
}

func (c *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

```

### Тесты

```go
package abstractFactory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("motorbike vehicle has %d wheels\n",
		motorbikeVehicle.NumWheels())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("struct assertion has failed")
	}

	t.Logf("sport motorbike has type %d\n", sportBike.GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
}

```
</details>

<details><summary> Prototype</summary>

### Prototype — избегание создания повторяющихся объектов.

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

</details>

********************************************
#### Структурные (Structural)
<details><summary> Composite</summary>
в процессе ...
</details>

<details><summary> Adapter</summary>
в процессе ...
</details>

<details><summary> Bridge</summary>
в процессе ...
</details>

<details><summary> Proxy</summary>
в процессе ...
</details>

<details><summary> Facade</summary>
в процессе ...
</details>

<details><summary> Flyweight</summary>
в процессе ...
</details>

<details><summary> Decorator</summary>
в процессе ...
</details>

********************************************
#### Поведенческие (Behavioral)
<details><summary> Strategy</summary>
в процессе ...
</details>

<details><summary> Chain of Responsibility</summary>
в процессе ...
</details>

<details><summary> Command</summary>
в процессе ...
</details>

<details><summary> Template</summary>
в процессе ...
</details>

<details><summary> Memento</summary>
в процессе ...
</details>

<details><summary> Interpreter</summary>
в процессе ...
</details>

<details><summary> Visitor</summary>
в процессе ...
</details>

<details><summary> State</summary>
в процессе ...
</details>

<details><summary> Mediator</summary>
в процессе ...
</details>

<details><summary> Observer</summary>
в процессе ...
</details>

********************************************
#### Конкурентные (Concurrency)
<details><summary> Barrier</summary>
в процессе ...
</details>

<details><summary> Future</summary>
в процессе ...
</details>

<details><summary> Pipeline</summary>
в процессе ...
</details>

<details><summary> Workers Pool</summary>
в процессе ...
</details>

<details><summary> Publish/Subscriber</summary>
в процессе ...
</details>



********************************************