package builder

type BuildProcess interface {
	SetSeats() BuildProcess
	SetDoors() BuildProcess
	SetWheels() BuildProcess
	SetColor() BuildProcess
	SetTransmition() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Seats        int
	Doors        int
	Wheels       int
	Color        string
	Transmission string
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(builder BuildProcess) {
	m.builder = builder
}

func (m *ManufacturingDirector) BuildVehicle() {
	m.builder.SetSeats().SetDoors().SetWheels().SetColor().SetTransmition()
}

type CarBuilder struct {
	vehicle VehicleProduct
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.vehicle.Seats = 5
	return c
}

func (c *CarBuilder) SetDoors() BuildProcess {
	c.vehicle.Doors = 4
	return c
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.vehicle.Wheels = 4
	return c
}

func (c *CarBuilder) SetTransmition() BuildProcess {
	c.vehicle.Transmission = "Automatic"
	return c
}

func (c *CarBuilder) SetColor() BuildProcess {
	c.vehicle.Color = "White"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicle
}

type CustomSportsCarBuilder struct {
	vehicle VehicleProduct
}

func (cs *CustomSportsCarBuilder) SetSeats(seats int) {
	cs.vehicle.Seats = seats
}

func (cs *CustomSportsCarBuilder) SetDoors(doors int) {
	cs.vehicle.Doors = doors
}

func (cs *CustomSportsCarBuilder) SetWheels(wheels int) {
	cs.vehicle.Wheels = wheels
}

func (cs *CustomSportsCarBuilder) SetTransmition(transmission string) {
	cs.vehicle.Transmission = transmission
}

func (cs *CustomSportsCarBuilder) SetColor(color string) {
	cs.vehicle.Color = color
}

func (cs *CustomSportsCarBuilder) GetVehicle() VehicleProduct {
	return cs.vehicle
}
