package builder

/*

Builder design pattern - reusing an algorithm to create many implementations of an interface

Description:
- The builder pattern helps us construct complex objects without directly instantiating their struct, or writing the
	logic they require

Objectives:
- A builder design pattern tries to:
	- Abstract complex creations so that object creation is separated from the thing that uses the object
	- Create an object step by step by filling its fields and creating the embedded objects
	- Reuse the object creation algorithm between many objects

Example:
- Vehicle manufacturing
- The builder pattern has been commonly described as the relationship between a director, a few Builders, and the
	products they build
- The process of creating a vehicle (the product) is more or less the same for every kind of vehicle -- choose vehicle
	type, assemble the structure, place the wheels, and place the seats
- If you think about it, you could build a car and a motorbike (two Builders) with this description.
- The director is represented by the `ManufacturingDirector` type in our example

After coding below...
Wrapping up the Builder design pattern:
- This design pattern helps us maintain an unpredictable number of products by using a common construction algorithm
	that is used by the director
- The construction process is always abstracted from the user of the product

*/

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct{
	builder BuildProcess
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
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

func (c *CarBuilder) GetVehicle() VehicleProduct {
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
	b.v.Seats = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4*2
	return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}

