package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Score        float32  `json:"score"`
	IsEligible   bool     `json:"is_eligible,omitempty"`
	TotalScore   float32  `json:"-"`
	ListSchedule []string `json:"list_schedule"`
}

type StudentList struct {
	Students []Student `json:"data"`
}

func main() {
	budi := Student{
		Name:       "Budi",
		Age:        15,
		Score:      80.5,
		IsEligible: true,
		TotalScore: 80.5,
		ListSchedule: []string{
			"Monday",
			"Tuesday",
			"Wednesday",
			"Thursday",
			"Friday",
		},
	}

	jsonData, err := json.Marshal(budi)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	students := StudentList{
		[]Student{
			{
				Name:       "Budi",
				Age:        15,
				Score:      80.5,
				IsEligible: true,
				TotalScore: 80.5,
				ListSchedule: []string{
					"Monday",
					"Tuesday",
					"Wednesday",
					"Thursday",
					"Friday",
				},
			},
			{
				Name:       "Budi 2",
				Age:        15,
				Score:      80.5,
				IsEligible: true,
				TotalScore: 80.5,
			},
		},
	}

	jsonListData, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonListData))

	userMap := make(map[string]Student)

	userMap["user_1"] = Student{
		Name:       "Budi",
		Age:        15,
		Score:      80.5,
		IsEligible: true,
		TotalScore: 80.5,
		ListSchedule: []string{
			"Monday",
			"Tuesday",
			"Wednesday",
			"Thursday",
			"Friday",
		},
	}
	jsonMapData, error := json.Marshal(userMap)
	if error != nil {
		panic(error)
	}

	jsonString := string(jsonMapData)
	fmt.Println(jsonString)
}
