package route

type Ship struct {
}

func (s *Ship) CalculatePrice(distance int64) int64 {
	return distance * 10000 / s.GetMaxPeople()
}

func (s *Ship) GetMaxPeople() int64 {
	return 100
}

func (s *Ship) GetName() string {
	return "Ship"
}
