package main

import "fmt"

type School interface {
	GetSchoolName(name string) string
	GetSchoolAddress(address string) string
}

type HighSchool struct{}

func (h *HighSchool) GetSchoolName(name string) string {
	return "Hello from High School " + name
}

func (h *HighSchool) GetSchoolAddress(address string) string {
	return "Hello from High School " + address
}

type JuniorHighSchool struct{}

func (j *JuniorHighSchool) GetSchoolName(name string) string {
	return "Junior High School " + name
}

func (j *JuniorHighSchool) GetSchoolAddress(address string) string {
	return "Junior High School " + address
}

type SchoolProvider interface {
	SetSchool() School
}

type LabHighSchool struct{}

func (l *LabHighSchool) SetSchool() School {
	return &HighSchool{}
}

type LabJuniorHighSchool struct{}

func (l *LabJuniorHighSchool) SetSchool() School {
	return &JuniorHighSchool{}
}

func main() {
	labHighSchool := &LabHighSchool{}
	labJuniorHighSchool := &LabJuniorHighSchool{}

	rawagunLabHighSchool := labHighSchool.SetSchool()
	rawagunLabJuniorHighSchool := labJuniorHighSchool.SetSchool()

	fmt.Println(rawagunLabHighSchool.GetSchoolName("Rawamangun Lab High School"))
	fmt.Println(rawagunLabHighSchool.GetSchoolAddress("Jl. Rawamangun"))

	fmt.Println(rawagunLabJuniorHighSchool.GetSchoolName("Rawamangun Lab Junior High School"))
	fmt.Println(rawagunLabJuniorHighSchool.GetSchoolAddress("Jl. Rawamangun"))
}
