package main

import "fmt"

// Logistics is a Creator
type Logistics interface {
	GetLogisticType() string
	CalculatePrice(day int) int
}

// SeaLogistics implements Logistics
// SeaLogistics is a Concrete Creator
type SeaLogistics struct{}

func (s SeaLogistics) GetLogisticType() string {
	return "Sea Logistics"
}

func (s SeaLogistics) CalculatePrice(day int) int {
	return day * 500000
}

// AirLogistics implements Logistics
// AirLogistics is a Concrete Creator
type AirLogistics struct{}

func (a AirLogistics) GetLogisticType() string {
	return "Air Logistics"
}

func (a AirLogistics) CalculatePrice(day int) int {
	return day * 1000000
}

type LogisticsProvider interface {
	SetLogistics() Logistics
}

type LautLepasLogistics struct{}

func (l LautLepasLogistics) SetLogistics() Logistics {
	return &SeaLogistics{}
}

type UdaraLepasLogistics struct{}

func (u UdaraLepasLogistics) SetLogistics() Logistics {
	return &AirLogistics{}
}

func main() {
	lautLepasLogistics := SeaLogistics{}
	fmt.Println(lautLepasLogistics.GetLogisticType())
	fmt.Println(lautLepasLogistics.CalculatePrice(10))

	udaraLepasLogistics := AirLogistics{}
	fmt.Println(udaraLepasLogistics.GetLogisticType())
	fmt.Println(udaraLepasLogistics.CalculatePrice(10))

	lll := LautLepasLogistics{}
	jakutLLL := lll.SetLogistics()
	fmt.Println(jakutLLL.GetLogisticType())
	fmt.Println(jakutLLL.CalculatePrice(20))

	ull := UdaraLepasLogistics{}
	jakutULL := ull.SetLogistics()
	fmt.Println(jakutULL.GetLogisticType())
	fmt.Println(jakutULL.CalculatePrice(20))
}
