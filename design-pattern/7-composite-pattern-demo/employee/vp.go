package employee

type VpEng struct {
	Subordinates []Employee
}

func (v *VpEng) GetEmplSalary() int {
	return 8000
}

func (v *VpEng) GetEmplTotalDivSalary() int {
	totalSalary := 0

	for _, sub := range v.Subordinates {
		totalSalary += sub.GetEmplSalary()
	}

	return totalSalary + v.GetEmplSalary()
}
