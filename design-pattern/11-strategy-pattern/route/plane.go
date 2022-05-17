package route

type Plane struct {
}

func (p *Plane) CalculatePrice(distance int64) int64 {
	return distance * 100000 / p.GetMaxPeople()
}

func (p *Plane) GetMaxPeople() int64 {
	return 250
}

func (p *Plane) GetName() string {
	return "Plane"
}
