package pizza

import "fmt"

type Pizza interface {
	GetPrice() int64
}

type PizzaBaseDecorator struct{}

func (p *PizzaBaseDecorator) GetPrice() int64 {
	return 1
}

type PizzaCheeseDecorator struct {
	PizzaBase Pizza
}

func (pc *PizzaCheeseDecorator) GetPrice() int64 {
	basePrice := pc.PizzaBase.GetPrice()

	pc.CheeseDesc()

	return basePrice + 10
}

func (pc *PizzaCheeseDecorator) CheeseDesc() {
	fmt.Println("Cheese")
}

type PizzaPepperoniDecorator struct {
	PizzaBase Pizza
}

func (pp *PizzaPepperoniDecorator) GetPrice() int64 {
	basePrice := pp.PizzaBase.GetPrice()

	pp.PepperoniDesc()

	return basePrice + 20
}

func (pp *PizzaPepperoniDecorator) PepperoniDesc() {
	fmt.Println("Pepperoni")
}
