package main

import "math"

type Calculator struct {
	acc float64
}

type opFunc func(float64, float64) float64

func (c *Calculator) Do(op opFunc, x float64) float64 {
	c.acc = op(c.acc, x)
	return c.acc
}

func Add(x, y float64) float64 {
	return x + y
}

func Sub(x, y float64) float64 {
	return x - y
}

func Mul(x, y float64) float64 {
	return x * y
}

func Sqrt(x, _ float64) float64 { // bad code
	return math.Sqrt(x)
}

func main() {
	c := new(Calculator)
	c.Do(Add, 10)
	c.Do(Sub, 2)
	c.Do(Mul, 3)
	c.Do(Sqrt, 0)
	println(c.acc)
}
