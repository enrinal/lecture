package military

type Military interface {
	GetSalary() int
	GetTotalDivSalary() int
}

type General struct {
	Subordinates []Military
}

func (g *General) GetSalary() int {
	return 10
}

func (g *General) GetTotalDivSalary() int {
	var total int
	for _, sub := range g.Subordinates {
		total += sub.GetTotalDivSalary()
	}
	return total + g.GetSalary()
}

type Major struct {
	Subordinates []Military
}

func (m *Major) GetSalary() int {
	return 8
}

func (m *Major) GetTotalDivSalary() int {
	var total int
	for _, sub := range m.Subordinates {
		total += sub.GetSalary()
	}
	return total + m.GetSalary()
}

type Soldier struct {
}

func (s *Soldier) GetSalary() int {
	return 1
}

func (s *Soldier) GetTotalDivSalary() int {
	return 0
}
