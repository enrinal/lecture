package main

import (
	"fmt"

	"lecture/design-pattern/9-decorative-pattern-demo/pizza"
)

func main() {
	p := &pizza.BasePizza{}

	fmt.Println(p.GetPizzaPrice())
	fmt.Println(p.GetPizzaTopping())

	pWithCheese := &pizza.PizzaCheese{
		PizzaWrappee: p,
	}
	fmt.Println("===========================")
	fmt.Println(pWithCheese.GetPizzaPrice())
	fmt.Println(pWithCheese.GetPizzaTopping())

	pWithCheeseMeat := &pizza.PizzaMeat{
		PizzaWrappee: pWithCheese,
	}
	fmt.Println("===========================")
	fmt.Println(pWithCheeseMeat.GetPizzaPrice())
	fmt.Println(pWithCheeseMeat.GetPizzaTopping())

	pWithMeat := &pizza.PizzaMeat{
		PizzaWrappee: p,
	}
	fmt.Println("===========================")
	fmt.Println(pWithMeat.GetPizzaPrice())
	fmt.Println(pWithMeat.GetPizzaTopping())
}
