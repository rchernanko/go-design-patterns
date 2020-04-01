package builder_test

import (
	"github.com/rchernanko/go-design-patterns/01_creational_patterns/builder"
	"testing"
)

func TestBuilderPattern_Car(t *testing.T) {
	manufacturingFactory := builder.ManufacturingDirector{}

	//First the car
	carBuilder := &builder.CarBuilder{}
	manufacturingFactory.SetBuilder(carBuilder)
	manufacturingFactory.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	//And then the bike
	bikeBuilder := &builder.BikeBuilder{}
	manufacturingFactory.SetBuilder(bikeBuilder)
	manufacturingFactory.Construct()

	bike := bikeBuilder.GetVehicle()
	bike.Seats = 1

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure on a bike must be 'Bike' and was %s\n", bike.Structure)
	}

	//And then the bus
	busBuilder := &builder.BusBuilder{}
	manufacturingFactory.SetBuilder(busBuilder)
	manufacturingFactory.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 8 {
		t.Errorf("Wheels on a bus must be 8 and they were %d\n", bus.Wheels)
	}

	if bus.Seats != 30 {
		t.Errorf("Seats on a bus must be 30 and they were %d\n", bus.Seats)
	}

	if bus.Structure != "Bus" {
		t.Errorf("Structure on a bus must be 'Bus' and was %s\n", bus.Structure)
	}
}