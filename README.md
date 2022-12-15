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

### Builder — повторное использование алгоритма для создания множества реализаций интерфейса

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

</details>

********************************************
#### Структурные (Structural)
<details><summary> Composite</summary>

### Composite — альтернатива наследования

### Описание

Шаблон проектирования Composite предпочитает композицию наследованию. Подход «композиция вместо наследования» был предметом дискуссий среди инженеров с девяностых годов. В общем, в Go нет наследования, потому что оно ему не нужно! В шаблоне проектирования Composite вы будете создавать иерархии и деревья объектов. Объекты имеют разные объекты со своими полями и методами внутри них. Этот подход очень мощный и решает многие проблемы наследования и множественного наследования.

Цель паттерна Composite состоит в том, чтобы избежать иерархического ада, когда сложность приложения может слишком сильно возрасти и это повлияет на ясность кода.

### Пример — пловец и акула

Типичная проблема наследования возникает, когда у вас есть объект, наследуемый от двух совершенно разных классов, между которыми нет абсолютно никакой связи. Представьте спортсмена, который тренируется и является пловцом с умением плавать:
* Athlete имеет метод Train().
* Swimmer имеет метод Swim().

Swimmer наследуется от Athlete, поэтому он наследует его метод Train и объявляет собственный метод Swim. У вас также может быть велосипедист, который также является спортсменом и объявляет метод Ride.
А теперь представьте себе Animal, например Shark, которая плавает, как и Swimmer. Ничего фантастического. Итак, как решить эту проблему? Акула не может быть пловцом, который еще и тренируется. Акулы не тренируются (насколько я знаю!).

Требования и критерии приемлемости:
* У нас должна быть структура Athlete с методом Train
* У нас должен быть Swimmer с методом Swim
* У нас должна быть структура Animal с методом Eat
* У нас должна быть структура Shark с методом Swim, который используется совместно со Swimmer

В Go мы можем использовать два типа композиции — прямую композицию и встраиваемую композицию. Сначала мы решим эту проблему, используя прямую композицию, которая имеет все необходимое в виде полей внутри структуры.


### Реализация с помощью нулевой инициализации
```go
package composite

import "fmt"

// Athlete

type Athlete struct {}

func (a *Athlete) Train() {
	fmt.Println("Training...")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

// Animal

type Animal struct {}

func (a *Animal) Eat() {
	fmt.Println("Eating...")
}

type Shark struct {
	Animal
	Swim func()
}

// Method for athlete and fish

func Swim() {
	fmt.Println("Swimming...")
```
### Реализация с помощью интерфейсов
```go
type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming...")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}
```
### Тесты
```go
package composite

import (
	"testing"
)

func TestAthleteA(t *testing.T) {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}

func TestAnimal(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestAthleteB(t *testing.T) {
	swimmer := CompositeSwimmerB{
		Trainer: &Athlete{},
		Swimmer: &SwimmerImpl{},
	}

	swimmer.Train()
	swimmer.Swim()
}

```
### Binary Tree compositions

Другой очень распространенный подход к шаблону Composite — это работа со структурами двоичного дерева. В двоичном дереве вам нужно хранить экземпляры самого себя в поле:
```go
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}
```
Это своего рода рекурсивная композиция, и из-за природы рекурсивности мы должны использовать указатели, чтобы компилятор знал, сколько памяти он должен зарезервировать для этой структуры. В нашей структуре Tree хранится объект LeafValue для каждого экземпляра и новое дерево в его полях Right и Left.
С помощью этой структуры мы могли бы создать объект и написать тест:
```go
func TestBinaryTree(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Left: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Right: &Tree{4, nil, nil},
	}
	right := root.Left.Right.LeafValue
	if right != 6 {
		t.Errorf("wrong result, must be 6, not %v", right)
	}
}
```
</details>

<details><summary> Adapter</summary>

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
</details>

<details><summary> Bridge</summary>

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