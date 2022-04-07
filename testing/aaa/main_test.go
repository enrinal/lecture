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

func TestCustomer_Pay(t *testing.T) {
	// arrange
	customerTests := []struct {
		name     string
		customer *Customer
		pay      float64
		want     float64
		wantErr  error
	}{
		{
			name: "success",
			customer: &Customer{
				Name: "Budi",
				Age:  20,
				Cash: float64(10000),
			},
			pay:  float64(5000),
			want: float64(5000),
		},
		{
			name: "insufficient funds",
			customer: &Customer{
				Name: "Budi",
				Age:  20,
				Cash: float64(10000),
			},
			pay:     float64(100000),
			want:    float64(10000),
			wantErr: errors.New("Insufficient funds"),
		},
	}
	for _, tt := range customerTests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			err := tt.customer.Pay(tt.pay)
			// assert
			assert.Equal(t, tt.want, tt.customer.Cash)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
