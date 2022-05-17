package route

type RouteStrategy interface {
	CalculatePrice(distance int64) int64
	GetMaxPeople() int64
	GetName() string
}
