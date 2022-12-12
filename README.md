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
В качестве примера объекта, будет уникальный счетчик, для которого мы должны убедиться, что существует только один такой экземпляр, счетчик будет сожержать количество вызовов во время исполнения программы. Неважно, сколько у нас экземпляров счетчика, все они будут считать одно и то же значение, и оно должно быть согласовано между экземплярами.

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

### Builder - повторное использование алгоритма для создания множества реализаций интерфейса.

### Описание

Шаблон Builder помогает нам создавать сложные объекты без непосредственного создания их структуры или написания необходимой логики. Представьте себе объект, который может иметь десятки полей, которые сами по себе являются более сложными структурами. Теперь представьте, что у вас есть много объектов с такими характеристиками. Здесь и пригодится Builder, чтобы не  писать логику для создания всех этих объектов.

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

<details><summary> Factory</summary>

</details>

<details><summary> Prototype</summary>
</details>

<details><summary> Abstract Factory</summary>
</details>

********************************************
#### Структурные (Structural)
<details><summary> Composite</summary>
</details>

<details><summary> Adapter</summary>
</details>

<details><summary> Bridge</summary>
</details>

<details><summary> Proxy</summary>
</details>

<details><summary> Facade</summary>
</details>

<details><summary> Flyweight</summary>
</details>

<details><summary> Decorator</summary>
</details>

********************************************
#### Поведенческие (Behavioral)
<details><summary> Strategy</summary>
</details>

<details><summary> Chain of Responsibility</summary>
</details>

<details><summary> Command</summary>
</details>

<details><summary> Template</summary>
</details>

<details><summary> Memento</summary>
</details>

<details><summary> Interpreter</summary>
</details>

<details><summary> Visitor</summary>
</details>

<details><summary> State</summary>
</details>

<details><summary> Mediator</summary>
</details>

<details><summary> Observer</summary>
</details>

********************************************
#### Конкурентные (Concurrency)
<details><summary> Barrier</summary>
</details>

<details><summary> Future</summary>
</details>

<details><summary> Pipeline</summary>
</details>

<details><summary> Workers Pool</summary>
</details>

<details><summary> Publish/Subscriber</summary>
</details>



********************************************