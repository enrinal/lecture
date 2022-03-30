package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Score      float32 `json:"score"`
	IsEligible bool    `json:"is_eligible,omitempty"`
}

func main() {
	budi := Student{
		Name:       "Budi",
		Age:        15,
		Score:      80.5,
		IsEligible: true,
	}

	jsonData, err := json.Marshal(budi) // marshal to json
	if err != nil {
		panic(err)
	}

	bunga := Student{}
	err = json.Unmarshal(jsonData, &bunga)
	if err != nil {
		panic(err)
	}
	bunga.Name = "Bunga"
	fmt.Println(bunga)

	byteJSONData := []byte(`{
    "name":"Rogu",
    "age":20,
    "score":80.5,
	"sex": "Men",
	"status": "ok"
	}`) // sex diabaikan karena tidak ada di struct student

	rogu := Student{}
	err = json.Unmarshal(byteJSONData, &rogu)
	if err != nil {
		panic(err)
	}

	fmt.Println(rogu)
}
