package main

import (
	"fmt"

	"lecture/design-pattern/7-composite-pattern/military"
)

func main() {
	soldier1 := &military.Soldier{}
	soldier2 := &military.Soldier{}
	soldier3 := &military.Soldier{}

	major1 := &military.Major{
		Subordinates: []military.Military{soldier1, soldier2},
	}

	major2 := &military.Major{
		Subordinates: []military.Military{soldier3},
	}

	general := &military.General{
		Subordinates: []military.Military{major1, major2},
	}

	total := general.GetTotalDivSalary()
	fmt.Println("Total salary:", total)
}
