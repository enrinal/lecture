package main

type Calculator struct {
	acc float64
}

const (
	OP_ADD = 1 << iota
	OP_SUB
	OP_MUL
)

func (c *Calculator) Do(op int, val float64) {
	switch op {
	case OP_ADD:
		c.acc += val
	case OP_SUB:
		c.acc -= val
	case OP_MUL:
		c.acc *= val
	default:
		panic("unknown op")
	}
}

func main() {
	c := &Calculator{}
	c.Do(OP_ADD, 1)
	c.Do(OP_ADD, 2)
	c.Do(OP_SUB, 1)
	c.Do(OP_MUL, 2)
	println(c.acc)
}
