package composite

import "fmt"

// Athlete

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training...")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

// Animal

type Animal struct{}

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
}

// Another method to use the Composite pattern for Athlete

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

// Binary tree Compositions

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}
