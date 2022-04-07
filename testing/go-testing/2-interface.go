package go_testing

import "math"

type Shape interface {
	Area() float64
}

type RectangleOne struct {
	Width  float64
	Height float64
}

type CircleOne struct {
	Radius float64
}

func (r *RectangleOne) Area() float64 {
	return r.Width * r.Height
}

func (c *CircleOne) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
