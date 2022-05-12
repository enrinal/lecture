package car

type europeCarBuilder struct {
	car Car
}

func (e *europeCarBuilder) buildDoors(numOfDoors int) {
	e.car.NumOfDoors = 2
}

func (e *europeCarBuilder) buildEngineType(engineType string) {
	e.car.EngineType = "electric"
}

func (e *europeCarBuilder) buildNumOfSeats(numOfSeats int) {
	e.car.NumOfSeats = 2
}

func (e *europeCarBuilder) getCar() Car {
	return e.car
}
