package main

import (
	"fmt"

	"lecture/design-pattern/10-state-pattern/light"
)

func main() {
	lightSwitch := light.NewLightSwitch()
	lightSwitch.Press()
	fmt.Println(lightSwitch.CanTurnOnLight())
	lightSwitch.Press()
	fmt.Println(lightSwitch.CanTurnOnLight())
}
