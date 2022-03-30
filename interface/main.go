package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

// Repository --> db --> interface StudentRepository
// Usecase --> logic --> interface StudentUsecase --> impl StudentRepository
// handler --> return http --> interface StudentHandler --> impl StudentUsecase

type StudentInterface interface {
	GetName() string
	GetAge() int
	GetScore() float64
	TranformScore(float64) string
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) GetScore() float64 {
	return s.Score
}

func (s *Student) TranformScore(score float64) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	budi := Student{
		Name:  "Budi",
		Age:   20,
		Score: 80.0,
	}

	fmt.Println(budi.GetName(), budi.GetAge(), budi.GetScore(), budi.TranformScore(budi.Score))
}
