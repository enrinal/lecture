package main

import (
	"fmt"

	"lecture/design-pattern/6-builder-pattern/car"
)

func main() {
	japanCarBuilder := car.NewCarBuilder("japan")
	europeCarBuilder := car.NewCarBuilder("europe")

	contractor := car.NewContractor(japanCarBuilder)
	familyCar := contractor.BuildFamilyCar()
	fmt.Printf("%+v\n", familyCar)
	japanSportCar := contractor.BuildSportCar()
	fmt.Printf("%+v\n", japanSportCar)

	contractor = car.NewContractor(europeCarBuilder)
	familyCar = contractor.BuildFamilyCar()
	fmt.Printf("%+v\n", familyCar)
	europeSportCar := contractor.BuildSportCar()
	fmt.Printf("%+v\n", europeSportCar)

}
