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