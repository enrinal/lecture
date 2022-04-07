package aaa

import "fmt"

type Customer struct {
	Name string
	Age  int
	Cash float64
}

func (c *Customer) Pay(amount float64) error {
	if c.Age < 18 {
		return fmt.Errorf("Customer is not old enough")
	}
	if c.Cash < amount {
		return fmt.Errorf("Insufficient funds")
	}
	c.Cash -= amount
	return nil
}
