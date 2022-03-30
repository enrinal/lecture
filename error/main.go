package main

import "fmt"

// only accept positive number
type Number struct {
	a int
	b int
}

type NumberInterface interface {
	Sum() (int, error)
	Div() (int, error)
	Sub() (int, error)
	MagicSum() (int, error)
}

var (
	errNegativeNumber = fmt.Errorf("negative number is not allowed")
	errDivByZero      = fmt.Errorf("divide by zero is not allowed")
)

func (n *Number) Sum() (int, error) {
	if n.a < 0 || n.b < 0 {
		return 0, errNegativeNumber
	}
	return n.a + n.b, nil
}

func (n *Number) Div() (int, error) {
	if n.a < 0 || n.b < 0 {
		return 0, errNegativeNumber
	}
	if n.b == 0 {
		return 0, errDivByZero
	}
	return n.a / n.b, nil
}

func (n *Number) Sub() (int, error) {
	if n.a < 0 || n.b < 0 {
		return 0, errNegativeNumber
	}
	return n.a - n.b, nil
}

func (n *Number) MagicSum() (int, error) {
	if n.a < 0 || n.b < 0 {
		return 0, errNegativeNumber
	}
	return n.a + n.b, nil
}

func main() {
	n := Number{a: 1, b: -1}
	resultSum, err := n.Sum()
	if err != nil {
		fmt.Println(err)
	}

	resultDiv, err := n.Div()
	if err != nil {
		fmt.Println(err)
	}

	resultSub, err := n.Sub()
	if err != nil {
		fmt.Println(err)
	}

	resultMagicSum, err := n.MagicSum()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultSum, resultDiv, resultSub, resultMagicSum)
}
