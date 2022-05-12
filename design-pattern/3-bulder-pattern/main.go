package main

type House struct {
	NumOfWindows    int
	NumOfDoors      int
	HasGarage       bool
	HasSwimmingPool bool
}
type HouseBuilder interface {
	buildWindow(numOfWindow int)
	buildDoor()
	buildGarage()
	buildSwimmingPool()
	getHouse() House
}

type contractor struct {
	builder HouseBuilder
}

func (c *contractor) BuildHouse() House {
	c.builder.buildWindow(5)
	c.builder.buildDoor()
	c.builder.buildGarage()
	c.builder.buildSwimmingPool()
	return c.builder.getHouse()
}
