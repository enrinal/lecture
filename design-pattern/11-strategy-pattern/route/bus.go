package route

type Bus struct{}

func (b *Bus) CalculatePrice(distance int64) int64 {
	return distance * 5000 / b.GetMaxPeople()
}

func (b *Bus) GetMaxPeople() int64 {
	return 60
}

func (b *Bus) GetName() string {
	return "Bus"
}
