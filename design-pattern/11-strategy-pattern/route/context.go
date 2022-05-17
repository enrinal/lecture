package route

import "fmt"

type Context struct {
	strategy RouteStrategy
}

func (c *Context) SetStrategy(s RouteStrategy) {
	c.strategy = s
}

func (c *Context) GetRoutePlan(distance int64) string {
	price := c.strategy.CalculatePrice(distance)
	maxPeople := c.strategy.GetMaxPeople()

	returnStr := fmt.Sprintf("%s: %dkm, %d people, %d rupiah", c.strategy.GetName(), distance, maxPeople, price)
	return returnStr
}
