package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	FirstName  string   `json:"firstName"`
	SecondName string   `json:"secondName"`
	Age        int      `json:"age"`
	Graduation []string `json:"graduation"`
}

func main() {
	// Read data from json data and unmarshal it into a struct
	person1Json := `{
	"firstName": "Leo",
	"secondName": "Suave",
	"age": 30,
	"graduation": ["Bachelor's Degree", "Master's Degree"]
	}`

	var person1 person

	if error := json.Unmarshal([]byte(person1Json), &person1); error != nil {
		log.Fatal(error)
	}

	fmt.Println(person1)

	// Read json data and unmarshal in map where the key is a string and the value is a person struct
	person2Json := `{
	"firstName": "Judite",
	"secondName": "Suave",
	"age": 6,
	"graduation": []
	}`

	var person2 person

	person2Map := make(map[string]person)

	if error := json.Unmarshal([]byte(person2Json), &person2); error != nil {
		log.Fatal(error)
	}
	person2Map["person"] = person2

	fmt.Println(person2Map)

	// Read data from json file and unmarshal it into a slice of person structs
	file, error := os.ReadFile("example.json")
	if error != nil {
		log.Fatal(error)
	}

	var person3 []person
	error = json.Unmarshal(file, &person3)
	if error != nil {
		log.Fatal(error)
	}

	for _, p := range person3 {
		fmt.Println(p)
	}
}
