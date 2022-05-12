package car

type japanCarBuilder struct {
	car Car
}

func (j *japanCarBuilder) buildDoors(numOfDoors int) {
	j.car.NumOfDoors = numOfDoors
}

func (j *japanCarBuilder) buildEngineType(engineType string) {
	if engineType == "gas" {
		j.car.EngineType = "gas"
	} else {
		j.car.EngineType = "diesel"
	}
}

func (j *japanCarBuilder) buildNumOfSeats(numOfSeats int) {
	j.car.NumOfSeats = numOfSeats
}

func (j *japanCarBuilder) getCar() Car {
	return j.car
}
