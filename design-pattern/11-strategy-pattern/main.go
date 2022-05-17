package main

import (
	"fmt"

	"lecture/design-pattern/11-strategy-pattern/route"
)

func main() {
	routeBus := route.Bus{}

	routePlan := route.Context{}
	routePlan.SetStrategy(&routeBus)

	fmt.Println(routePlan.GetRoutePlan(100))

	routeShip := route.Ship{}
	routePlan.SetStrategy(&routeShip)

	fmt.Println(routePlan.GetRoutePlan(100))

	routePlane := route.Plane{}
	routePlan.SetStrategy(&routePlane)

	fmt.Println(routePlan.GetRoutePlan(100))

}
