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
