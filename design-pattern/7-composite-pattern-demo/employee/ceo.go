package employee

type CEO struct {
	Subordinates []Employee
}

func (c *CEO) GetEmplSalary() int {
	return 10000
}
func (c *CEO) GetEmplTotalDivSalary() int {
	totalSalary := 0

	for _, sub := range c.Subordinates {
		totalSalary += sub.GetEmplTotalDivSalary()
	}

	return totalSalary + c.GetEmplSalary()
}
