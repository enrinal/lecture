package car

type Car struct {
	NumOfDoors int
	EngineType string
	NumOfSeats int
}

type CarBuilder interface {
	buildDoors(numOfDoors int)
	buildEngineType(engineType string)
	buildNumOfSeats(numOfSeats int)
	getCar() Car
}

func NewCarBuilder(builderType string) CarBuilder {
	if builderType == "japan" {
		return &japanCarBuilder{}
	}

	if builderType == "europe" {
		return &europeCarBuilder{}
	}

	return nil
}
