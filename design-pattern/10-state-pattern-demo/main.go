package main

import (
	"fmt"

	"lecture/design-pattern/10-state-pattern-demo/light"
)

func main() {
	ls := light.NewLightSwitch()

	fmt.Println("can turn on light:", ls.CanLightTurnOn())
	fmt.Println("light condition:", ls.GetLightCondition())

	fmt.Println("==============================")
	ls.LightPress()
	fmt.Println("can turn on light:", ls.CanLightTurnOn())
	fmt.Println("light condition:", ls.GetLightCondition())

	fmt.Println("==============================")
	ls.LightPress()
	fmt.Println("can turn on light:", ls.CanLightTurnOn())
	fmt.Println("light condition:", ls.GetLightCondition())

	fmt.Println("==============================")
	ls.LightPress()
	fmt.Println("can turn on light:", ls.CanLightTurnOn())
	fmt.Println("light condition:", ls.GetLightCondition())
}
