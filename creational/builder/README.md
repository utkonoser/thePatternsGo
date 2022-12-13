### Builder — повторное использование алгоритма для создания множества реализаций интерфейса.

### Описание

Шаблон Builder помогает нам создавать сложные объекты без непосредственного создания их структуры или написания необходимой логики. Представьте себе объект, который может иметь десятки полей, которые сами по себе являются более сложными структурами. Теперь представьте, что у вас есть много объектов с такими характеристиками. Здесь и пригодится Builder, чтобы не писать логику для создания всех этих объектов.

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