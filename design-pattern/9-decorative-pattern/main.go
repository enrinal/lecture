package main

import (
	"fmt"

	"lecture/design-pattern/9-decorative-pattern/pizza"
)

func main() {

	pizzaBase := &pizza.PizzaBaseDecorator{}

	pizzaCheese := &pizza.PizzaCheeseDecorator{
		PizzaBase: pizzaBase,
	}

	fmt.Println(pizzaCheese.GetPrice())

	pizzaCheesePepperoni := &pizza.PizzaPepperoniDecorator{
		PizzaBase: pizzaCheese,
	}

	fmt.Println(pizzaCheesePepperoni.GetPrice())
}
