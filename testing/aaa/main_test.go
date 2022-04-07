package aaa

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomer_Pay_LimitAge(t *testing.T) {
	// arrange
	c := Customer{
		Name: "John",
		Age:  16,
		Cash: float64(200000),
	}
	pay := float64(100000)

	// act
	err := c.Pay(pay)

	// assert
	assert.Equal(t, float64(200000), c.Cash)
	assert.Error(t, errors.New("Customer is not old enough"), err)
}

func TestCustomer_Pay_InsufficientFunds(t *testing.T) {
	// arrange
	budi := &Customer{
		Name: "Budi",
		Age:  20,
		Cash: float64(10000),
	}
	pay := float64(100000)

	// act
	err := budi.Pay(pay)

	// assert
	assert.Equal(t, float64(10000), budi.Cash)
	assert.Error(t, errors.New("Insufficient funds"), err)
}

func TestCustomer_Pay_Success(t *testing.T) {
	// arrange
	budi := &Customer{
		Name: "Budi",
		Age:  20,
		Cash: float64(10000),
	}
	pay := float64(5000)

	// act
	err := budi.Pay(pay)

	// assert
	assert.Equal(t, float64(5000), budi.Cash)
	assert.NoError(t, err)
}
