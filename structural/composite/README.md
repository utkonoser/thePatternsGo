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