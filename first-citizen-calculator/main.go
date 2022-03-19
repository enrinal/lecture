package main

import "math"

type Calculator struct {
	acc float64
}

func (c *Calculator) Do(op func(float64) float64) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func Add(x float64) func(float64) float64 {
	return func(y float64) float64 {
		return y + x
	}
}

func Sub(x float64) func(float64) float64 {
	return func(y float64) float64 {
		return y - x
	}
}

func Mul(x float64) func(float64) float64 {
	return func(y float64) float64 {
		return y * x
	}
}

func Sqrt() func(float64) float64 {
	return math.Sqrt
}

func main() {
	c := &Calculator{}
	c.Do(Add(10))
	c.Do(Sub(2))
	c.Do(Mul(3))
	c.Do(Sqrt())

	println(c.acc)
}
