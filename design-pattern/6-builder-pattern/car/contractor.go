package car

type Contractor struct {
	carBuilder CarBuilder
}

func NewContractor(builder CarBuilder) *Contractor {
	return &Contractor{
		carBuilder: builder,
	}
}

func (c *Contractor) BuildFamilyCar() Car {
	c.carBuilder.buildDoors(4)
	c.carBuilder.buildEngineType("gas")
	c.carBuilder.buildNumOfSeats(7)
	return c.carBuilder.getCar()
}

func (c *Contractor) BuildSportCar() Car {
	c.carBuilder.buildDoors(2)
	c.carBuilder.buildEngineType("gas")
	c.carBuilder.buildNumOfSeats(2)
	return c.carBuilder.getCar()
}

func (c *Contractor) BuildLuxuryCar() Car {
	c.carBuilder.buildDoors(2)
	c.carBuilder.buildEngineType("electric")
	c.carBuilder.buildNumOfSeats(2)
	return c.carBuilder.getCar()
}
