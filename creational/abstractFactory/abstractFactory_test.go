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
