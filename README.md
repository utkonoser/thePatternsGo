## The Patterns - Go
Паттерны, реализованные в Go с примерами для обучения.

Репозиторий представляет собой набор реализаций различных паттернов с открытым исходным кодом.
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

### Proxy — оборачивание объекта для сокрытия характеристик

### Описание

Шаблон Proxy обычно оборачивает объект, чтобы скрыть некоторые его характеристики. Эти характеристики могут заключаться в том, что это удаленный объект (remote proxy), очень тяжелый объект, такой как дамп терабайтной базы данных (virtual proxy), или объект с ограниченным доступом (protection proxy).

Возможностей паттерна Proxy много, но в целом все они пытаются обеспечить одни и те же следующие функции:
* Скрыть объект за прокси-сервером для того, чтобы возможные функции можно было скрыть или ограничить
* Обеспечить новый уровень абстракции, с которым легко работать и можно легко изменить

### Пример

В примере мы собираемся создать удаленный прокси, который будет кэшировать объекты перед доступом к базе данных. Давайте представим, что у нас есть база данных со многими пользователями, но вместо того, чтобы обращаться к базе данных каждый раз, когда нам нужна информация о пользователе, у нас будет стек пользователей в порядке поступления (FIFO) в шаблоне Proxy.

Требования и критерии приемлемости:
* Весь доступ к базе данных пользователей будет осуществляться через тип Proxy
* Стек из n последних пользователей будет храниться в Proxy
* Если пользователь уже существует в стеке, запроса в базу данных не будет, вернется кешируемое значение
* Если запрошенный пользователь не существует в стеке, будет сделан запрос в базу данных, если стек полон, то удалим самого старого пользователя в стеке, далее сохраним нового пользователя и вернем его

### Реализация
```go
package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found\n", id)
}

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}
	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)
	fmt.Println("returning user from database")
	u.DidLastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

```

### Тесты

```go
package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCache:    UserList{},
		StackCapacity: 2,
	}

	knownIDs := [3]int32{someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("after one successful search empty cache, the size of it must be one")
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("no user can be returned from empty cache")
		}
	})

	t.Run("FindUser - one user, ask fo the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("cache must not grow if we asked for an object that is stored on it")
		}
		if !proxy.DidLastSearchUsedCache {
			t.Error("the user should have been returned from the cache")
		}
	})

	user1, err := proxy.FindUser(knownIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	user2, _ := proxy.FindUser(knownIDs[1])
	if proxy.DidLastSearchUsedCache {
		t.Error("the user wasn't stored on the proxy cache yet")
	}

	user3, _ := proxy.FindUser(knownIDs[2])
	if proxy.DidLastSearchUsedCache {
		t.Error("the user wasn't stored on the proxy cache yet")
	}

	for i := 0; i < len(proxy.StackCache); i++ {
		if proxy.StackCache[i].ID == user1.ID {
			t.Error("user that should be gone was found")
		}
	}

	if len(proxy.StackCache) != 2 {
		t.Error("after inserting 3 users the cache should not grow more than to two")
	}

	for _, v := range proxy.StackCache {
		if v != user2 && v != user3 {
			t.Error("a non expected user was found on the cache")
		}
	}

}

```
</details>

<details><summary> Decorator</summary>

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
	t .Log(pizzaResult)
}

```
</details>

<details><summary> Facade</summary>

### Facade — создание библиотеки

### Описание

Если прочитать про паттерн Proxy, то можно узнать, что это способ обернуть тип, чтобы скрыть от пользователя некоторые его сложные особенности. А если представить, что можно сгруппировать множество прокси в одной точке, например, в файле или библиотеке. Это может быть паттерн Facade. Facade в архитектурном смысле – это фасадная стена, скрывающая помещения и коридоры здания. Он защищает своих обитателей от холода и дождя и обеспечивает им уединение.
Шаблон проектирования Facade делает то же самое, но в нашем коде. Он защищает код от нежелательного доступа, упорядочивает некоторые вызовы и скрывает сложные области от пользователя.

Самый яркий пример паттерна Facade — это библиотека, где кто-то должен предоставить разработчику некоторые методы для выполнения определенных действий в дружественной манере. Таким образом, если разработчику нужно использовать вашу библиотеку, ему не нужно знать все внутренние задачи, чтобы получить желаемый результат.

Итак, вы паттерн проектирования Facade используется в следующих сценариях:
* Если нужно уменьшить сложность некоторых частей кода, то можно эту сложность скрыть с помощью Facade, предоставляя более простой в использовании метод
* Если нужно сгруппировать взаимосвязанные действия в одном месте
* Когда нужно создать библиотеку, чтобы другие могли использовать ваши продукты, не беспокоясь о том, как все это работает

### Пример — HTTP REST API для OpenWeatherMaps

В качестве примера попробуем написать часть библиотеки, которая обращается к сервису OpenWeatherMaps. Если вы не знакомы с сервисом OpenWeatherMap, это HTTP-сервис, который предоставляет вам оперативную информацию о погоде. HTTP REST API очень прост в использовании и станет хорошим примером того, как реализовать паттерн Facade для сокрытия сложности сетевых подключений за службой REST.

API OpenWeatherMap предоставляет много информации, поэтому мы сосредоточимся на получении данных о погоде в реальном времени в одном городе в каком-либо географическом месте, используя его значения широты и долготы. Ниже приведены требования и критерии приемлемости для этого шаблона проектирования:
* Нужно реализовать единый тип для доступа к данным. Вся информация, полученная из сервиса OpenWeatherMap, будет проходить через него
* Создать способ получения данных о погоде для какого-либо города какой-либо страны
* Создать способ получения данных о погоде для определенной широты и долготы
* Снаружи пакета должно быть видно не все, только самое важное, остальное должно быть скрыто (включая все данные, связанные с соединением)

### Реализация
```go
package facade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

type CurrentWeatherData struct {
	APIkey string
}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%v,%v&APPID=%s", lat, lon, c.APIkey))
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&APPID=%s", city, countryCode, c.APIkey))
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		byt, errMsg := io.ReadAll(resp.Body)

		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n", resp.StatusCode, errMsg)
		return
	}
	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()
	return
}

// getMockData - mock data for our example
func getMockData() io.Reader {
	response := `{
"coord":{"lon":-3.7,"lat":40.42},"weather":
[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],
"base":"stations","main":{"temp":303.56,"pressure":1016.46,
"humidity":26.8,"temp_min":300.95,"temp_max":305.93},"wind":{"speed":3.17,"deg":151.001},
"rain":{"3h":0.0075},"clouds":{"all":68},"dt":1471295823,"sys":{"type":3,"id":1442829648,
"message":0.0278,"country":"ES","sunrise":1471238808,"sunset":1471288232},"id":3117735,
"name":"Madrid","cod":200}`
	r := bytes.NewReader([]byte(response))
	return r
}

// Weather struct from http://openweathermap.org/current#current_JSON.
type Weather struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}
```

### Тесты

```go
package facade

import (
	"fmt"
	"testing"
)

// test with mock data
func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
}

// if there is api then use this test
func TestWithApi(t *testing.T) {
	weatherMap := CurrentWeatherData{"*Apikey"}
	weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Temperature in Madrid is %f celsius\n", weather.Main.Temp-273.15)
}

```
</details>

<details><summary> Flyweight</summary>

### Flyweight — паттерн Приспособленец

### Описание

Паттерн проектирования Flyweight. Он очень часто используется в компьютерной графике и индустрии видеоигр, но не так часто в корпоративных приложениях.
Flyweight — это паттерн, который позволяет разделить состояние тяжелого объекта между многими экземплярами одного типа. Представьте, что вам нужно создать и хранить слишком много принципиально одинаковых объектов, занимающих довольно много памяти. У вас быстро закончится память. Эта проблема может быть легко решена с помощью паттерна Flyweight с дополнительной помощью паттерна Factory. Благодаря шаблону Flyweight мы можем разделить все возможные состояния объектов в одном общем объекте и, таким образом, свести к минимуму создание объектов, используя указатели на уже созданные объекты.

### Пример — букмекерская контора
Чтобы привести пример, мы собираемся смоделировать то, что можно найти на странице букмекерской конторы. Представьте себе финальный матч чемпионата Европы, который смотрят миллионы людей по всему континенту. Теперь представьте, что у нас есть веб-страница для ставок, на которой мы публикуем историческую информацию о каждой команде в Европе. Это очень много информации, которая обычно хранится в какой-то распределенной базе данных, и у каждой команды буквально мегабайты информации о своих игроках, матчах, чемпионатах и так далее. Если миллион пользователей получат доступ к информации о команде и для каждого пользователя, запрашивающего исторические данные, будет создан новый экземпляр информации, у нас в мгновение ока закончится память. Можно попробовать Proxy паттерн, где можно кэшировать n самых последних поисков, чтобы ускорить запросы, но если мы будем возвращать клон для каждой команды, нам все равно будет не хватать памяти.
Вместо этого мы будем хранить информацию о каждой команде только один раз, и будем предоставлять ссылки на них пользователям. Таким образом, если мы столкнемся с миллионом пользователей, пытающихся получить доступ к информации о матче, у нас фактически будет просто две команды в памяти с миллионом указателей на одно и то же направление памяти.

Критерии приемлемости паттерна Flyweight всегда должны уменьшать объем используемой памяти и должны быть сосредоточены в первую очередь на этой цели:
*  Мы создадим структуру Team с некоторой базовой информацией, такой как название команды, игроки, прошлые результаты матчей, изображение их эмблемы
* Мы должны обеспечить правильное создание команды и отсутствие дубликатов
* При создании одной и той же команды дважды у нас должно быть два указателя, указывающих на один и тот же адрес памяти

### Реализация
```go
package flyweight

import "time"

type Team struct {
	Id             uint64
	Name           int
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Lear          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			Id:   2,
			Name: TEAM_B,
		}
	default:
		return Team{
			Id:   1,
			Name: TEAM_A,
		}
	}
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{createdTeams: make(map[int]*Team)}
}

```
### Тесты
```go
package flyweight

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("the pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}
	if teamA1 != teamA2 {
		t.Error("TEAM_A pointers weren't the same")
	}
	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	factory := NewTeamFactory()
	teams := make([]*Team, 500000*2)
	for i := 0; i < 500000; i++ {
		teams[i] = factory.GetTeam(TEAM_A)
	}
	for i := 500000; i < 2*500000; i++ {
		teams[i] = factory.GetTeam(TEAM_B)
	}
	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("The number of objects created was not 2: %d\n", factory.GetNumberOfObjects())
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}

```
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
<details><summary> Concurrent Singleton</summary>

### Concurrent Singleton — используя мьютексы и каналы

### Описание

В Creational паттернах есть паттерн Singleton — это некая структура или переменная, которая существует в коде только один раз. Весь доступ к этой структуре должен осуществляться с использованием описанного паттерна, но на самом деле он не безопасен с параллельной точки зрения.
Concurrent Singleton будет описан с учетом параллелизма.

### Пример — уникальный счетчик с помощью каналов и мьютексов
Чтобы ограничить одновременный доступ к экземпляру Singleton, только одна горутина сможет получить к нему доступ. Мы получим доступ к нему, используя каналы — первый для добавления единицы в счетчик, второй для получения текущего счетчика и третий для остановки горутины.
Мы добавим единицу в счетчик 10 000 раз, используя 10 000 различных горутин, запущенных из двух разных экземпляров Singleton. Затем мы введем цикл для проверки количества Singleton до тех пор, пока оно не станет равным 5000, и напишем значение счетчика перед запуском цикла.
Как только счетчик достигнет 5000, цикл завершится и закроет запущенную горутину.

### Реализация с помощью каналов

```go
package concurrentSingleton

import "sync"

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
```

### Реализация с помощью мьютексов
```go
type singleton2 struct {
	count int
	sync.RWMutex
}

var instance2 singleton2

func GetInstance2() *singleton2 {
	return &instance2
}
func (s *singleton2) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}
func (s *singleton2) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
```

### Тесты
```go
package concurrentSingleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()
}

func TestStartInstanceMutex(t *testing.T) {
	singleton := GetInstance2()
	singleton2 := GetInstance2()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
}
```
</details>

<details><summary> Barrier</summary>

### Barrier — ожидание всех горутин и вывод одного общего результата

### Описание

Представьте ситуацию, когда у нас есть приложение микросервисов, в котором один микросервис должна составить свой ответ путем слияния ответов трех других микросервисов. Здесь и поможет паттерн Barrier.
Barrier может быть сервисом, который будет блокировать свой ответ до тех пор, пока он не будет составлен из результатов, возвращаемых одной или несколькими различными горутинами (или сервисами).
Как следует из названия, паттерн Barrier пытается остановить выполнение, чтобы оно не завершилось до тех пор, пока все нужные результаты не будут готовы.

### Пример — HTTP GET aggregator

Для примера мы собираемся написать очень типичную ситуацию в приложении микросервисов — приложении, выполняющего два вызова HTTP GET и объединяющего их в один ответ, который будет напечатан на консоли.
Наше небольшое приложение должно выполнять каждый запрос в горутине и выводить результат на консоль, если оба ответа верны. Если какой-либо из них возвращает ошибку, мы печатаем только ошибку.

Требования и критерии приемлемости:
* Выведите на консоль объединенный результат двух вызовов URL-адресов `http://httpbin.org/headers` и `http://httpbin.org/User-Agent`. Это пара общедоступных конечных точек, которые отвечают данными о входящих соединениях
* Если какой-либо из вызовов терпит неудачу, он не должен печатать никакого результата — только сообщение об ошибке (или сообщения об ошибках, если оба вызова не удались)
* Вывод должен быть напечатан как составной результат после завершения обоих вызовов. Это означает, что мы не можем вывести результат одного вызова, а затем другого

### Реализация
```go
package barrier

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

```

### Тесты

```go
package barrier

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || strings.Contains(result, "User-Agent") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out
	return temp
}

```
</details>

<details><summary> Future</summary>

### Future — реализация принципа 'fire-and-forget'

### Описание

Паттерн проектирования Future (также называемый Promise) — это быстрый и простой способ создания конкурентных структур для асинхронного программирования. Идея здесь состоит в том, чтобы реализовать принцип «fire-and-forget», который обрабатывает все возможные результаты действия.
Короче говоря, мы определим каждое возможное поведение действия перед его выполнением в разных горутинах. Здесь интересно то, что мы можем запустить новый Future внутри Future и встроить столько Future, сколько захотим, в одну и ту же горутину (или в новые).

С помощью паттерна Future мы можем запускать множество новых горутин, каждая из которых имеет действие и собственный обработчик. Это позволяет нам сделать следующее:
* Делегировать обработчик действий другой горутине
* Стекировать между собой множество асинхронных вызовов (асинхронный вызов, который в своих результатах вызывает другой асинхронный вызов)

### Пример — простой асинхронный requester
В этом примере у нас будет метод, который возвращает строку или ошибку, но мы хотим выполнить все конкурентно. Используя канал, мы можем запустить новую горутину и обработать входящий результат из канала.
Но в этом случае нам придется обрабатывать результат (строку или ошибку), а этого мы не хотим. Вместо этого мы определим, что делать в случае успеха и что делать в случае ошибки.

Требования и критерии приемлемости:
* Делегировать выполнение функции другой горутине
* Функция вернет string (maybe) или error
* Обработчики должны быть уже определены до выполнения функции
* Дизайн должен быть многоразовым

### Реализация
```go
package Future

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}
```

### Тесты
```go
package Future

import (
	"errors"
	"sync"
	"testing"
)

func TestStringOrError(t *testing.T) {
	future := &MaybeString{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})
	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "", errors.New("error occurred")
		})
		wg.Wait()
	})
}
```
</details>

<details><summary> Pipeline</summary>

### Pipeline — конвейер передачи данных

### Описание

Pipeline — это паттерн, предназначенный для соединения горутин
и каналов, так что выходные данные одной горутины становятся входными данными для другой горутины, а для передачи данных используются каналы.
Одним из преимуществ использования конвейеров является наличие постоянного потока данных, так что никакие горутины и каналы не должны ожидать, пока завершится все остальное, чтобы можно было начать выполнение. Кроме того, мы используем меньше переменных и, следовательно, меньше памяти, потому что не приходится сохранять все данные в виде переменных. Наконец, использование конвейеров упрощает разработку программ и делает их удобнее для поддержки.

Возможности паттерна Pipeline:
* Можно создать параллельную структуру многошагового алгоритма
* Можно использовать параллелизм многоядерных машин, разложив алгоритм на разные горутины

### Пример — конвейер математических операций

Мы собираемся сгенерировать список чисел, начиная с 1 и заканчивая некоторым произвольным числом N. Затем мы возьмем каждое число, возведем его в степень 2 и суммируем полученные числа с уникальным результатом. Итак, если `N = 3`, наш список будет `[1,2,3]`. После включения их в 2 наш список становится `[1,4,9]`. Если мы суммируем полученный список, результирующее значение равно 14.

План реализации:
* Нужно создать список от 1 до N, где N может быть любым целым числом
* Взять каждое число из этого сгенерированного списка и возвести его в степень 2
* Суммировать каждое полученное число в окончательный результат и вернуть его

### Реализация
```go
package pipeline

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)
	result := <-thirdCh
	return result
}

// LaunchPipeline function doesn't need to allocate every channel, and can be rewritten like this:

func LaunchPipelineSecondVar(amount int) int {
	return <-sum(power(generator(amount)))
}
func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()
	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

```
### Тесты
```go
package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}

func TestLaunchPipelineSecondVar(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipelineSecondVar(test[0])
		if res != test[1] {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}

```
</details>

<details><summary> Workers Pool</summary>

### Workers pool — пул обработчиков

### Описание

Workers pool — это множество потоков, предназначенных для обработки назначаемых им заданий. Веб-сервер Apache и Go-пакет net/http работают приблизительно так: основной процесс принимает все входящие запросы, которые затем перенаправляются рабочим процессам для обработки. Как только рабочий процесс завершает свою работу, он готов к обслуживанию нового клиента.
Однако здесь есть главное различие: пул обработчиков использует не потоки,
а горутины. Кроме того, потоки обычно не умирают после обработки запросов, потому что затраты на завершение потока и создание нового слишком высоки, тогда как горутина прекращает существовать после завершения работы. Пулы обработчиков в Go реализованы с помощью буферизованных каналов, поскольку они позволяют ограничить число одновременно выполняемых горутин.

Создание Workers pool связано с управлением ресурсами: ЦП, ОЗУ, временем, соединениями и так далее. Паттерн проектирования Workers pool помогает сделать следующее:
* Контролировать доступ к общим ресурсам
* Создавать ограниченное количество горутин для каждого приложения
* Обеспечить больше возможностей параллелизма для других конкурентных структур

### Пример — pool of pipelines
В примере мы запустим ограниченное количество Pipeline, чтобы планировщик Go мог попытаться обрабатывать запросы параллельно.
Идея здесь состоит в том, чтобы контролировать количество горутин, изящно останавливать их, когда приложение завершает работу, и максимизировать параллелизм. В примере мы будем передавать строки, к которым будем добавлять данные и префиксы.

### Реализация
```go
package workersPool

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// workers pipeline

type Request struct {
	Data    any
	Handler RequestHandler
}

type RequestHandler func(any)

func NewStringRequest(data string, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: data, Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}

// worker

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type PrefixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (w *PrefixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercaseStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, uppercaseStringWithSuffix))
		}
	}()
}

// dispatcher

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(r Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{inCh: make(chan Request, b)}
}
```

### Тесты
```go
package workersPool

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkersPool(t *testing.T) {
	bufferSize := 100
	dispatcher := NewDispatcher(bufferSize)
	workers := 3
	for i := 1; i <= workers; i++ {
		var w WorkerLauncher = &PrefixSuffixWorker{
			id:      i,
			prefixS: fmt.Sprintf("Worker id: %d -> ", i),
			suffixS: " World",
		}
		dispatcher.LaunchWorker(w)
	}
	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()
}

```

</details>

<details><summary> Publish/Subscriber</summary>
в процессе ...
</details>



********************************************