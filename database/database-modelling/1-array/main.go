package main

import "fmt"

type StudentRow struct {
	ID          int // ID is the primary key
	NIM         string
	StudentName string
	Class       string
	SPP         int
	NoHP        string
}

type StudentDB []StudentRow // StudentDB is a table of students

func (s *StudentDB) Insert(name string, nim string, class string, spp int, nohp string) {
	// O(1)
	*s = append(*s, StudentRow{
		ID:          len(*s) + 1,
		NIM:         nim,
		StudentName: name,
		Class:       class,
		SPP:         spp,
		NoHP:        nohp,
	})
}

func (s *StudentDB) Where(id int) *StudentRow {
	// O(n)
	for _, student := range *s {
		if student.ID == id {
			return &student
		}
	}
	return nil
}

func main() {
	var students StudentDB
	students.Insert("John", "123456781", "B-01", 100000, "08123456781")
	students.Insert("Jane", "123456782", "C-04", 300000, "08123456782")
	students.Insert("Gina", "123456783", "A-03", 230000, "08123456783")
	students.Insert("Vina", "123456784", "C-01", 240000, "08123456784")

	fmt.Println(students.Where(3))
}
