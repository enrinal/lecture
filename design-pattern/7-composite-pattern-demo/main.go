package main

import (
	"fmt"

	"lecture/design-pattern/7-composite-pattern-demo/employee"
)

func main() {
	krocoSatu := &employee.Kroco{}
	krocoDua := &employee.Kroco{}
	krocoTiga := &employee.Kroco{}
	krocoEmpat := &employee.Kroco{}

	vpEngSatu := &employee.VpEng{
		Subordinates: []employee.Employee{krocoSatu, krocoDua, krocoTiga},
	}

	vpEngDua := &employee.VpEng{
		Subordinates: []employee.Employee{krocoEmpat},
	}

	ceo := &employee.CEO{
		Subordinates: []employee.Employee{vpEngSatu, vpEngDua},
	}

	//

	fmt.Println("Total salary div vp eng satu: ", vpEngSatu.GetEmplTotalDivSalary()) // 23000
	fmt.Println("Total salary div vp eng dua: ", vpEngDua.GetEmplTotalDivSalary())   // 13000

	fmt.Println("Total salary company:", ceo.GetEmplTotalDivSalary()) // 46000
}
