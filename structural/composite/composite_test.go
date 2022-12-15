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
