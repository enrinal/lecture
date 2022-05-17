package pizza

type PizzaCheese struct {
	PizzaWrappee Pizza
}

func (p *PizzaCheese) GetPizzaPrice() int64 {
	return p.PizzaWrappee.GetPizzaPrice() + 10
}

func (p *PizzaCheese) GetPizzaTopping() string {
	return p.PizzaWrappee.GetPizzaTopping() + " + Cheese"
}
