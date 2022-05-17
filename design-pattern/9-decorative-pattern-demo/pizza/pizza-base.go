package pizza

type BasePizza struct {
	PizzaWrappee Pizza
}

func (b *BasePizza) GetPizzaPrice() int64 {
	return 1
}

func (b *BasePizza) GetPizzaTopping() string {
	return "Create a pizza base"
}
