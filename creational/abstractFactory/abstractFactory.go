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
