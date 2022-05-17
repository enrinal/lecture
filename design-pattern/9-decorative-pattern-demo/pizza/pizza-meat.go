package pizza

type PizzaMeat struct {
	PizzaWrappee Pizza
}

func (p *PizzaMeat) GetPizzaPrice() int64 {
	return p.PizzaWrappee.GetPizzaPrice() + 20
}

func (p *PizzaMeat) GetPizzaTopping() string {
	return p.PizzaWrappee.GetPizzaTopping() + " + Meat"
}
