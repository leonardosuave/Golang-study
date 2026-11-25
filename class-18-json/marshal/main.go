package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// declar a struct and how would be converted to json
type person struct {
	FirstName  string   `json:"firstName"`
	SecondName string   `json:"secondName"`
	Age        int      `json:"age"`
	Graduation []string `json:"graduation"`
}

func main() {
	// json.Marshal() -> convert map or struct to json
	// json.Unmarshal() -> convert json to map or struct

	person1 := person{
		FirstName:  "John",
		SecondName: "Doe",
		Age:        30,
		Graduation: []string{"Bachelor's Degree", "Master's Degree"},
	}

	person1JSON, erro := json.Marshal(person1)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(person1JSON) // Will print bytes slice and not a json
	fmt.Println(string(person1JSON)) // Will print a json by a direct conversion of bytes slice to string
	fmt.Println(bytes.NewBuffer(person1JSON)) // Will print a json by a conversion of bytes slice to string using bytes.NewBuffer()

	// Map where person is the value of the map and the key is a string and the person struct is the value of the map
	person2 := map[string]person{
		"person": {"Jane", "Doe", 25, []string{"Bachelor's Degree"}},
	}

	person2JSON, erro := json.Marshal(person2["person"])
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(person2JSON) // Will print bytes slice and not a json
	fmt.Println(bytes.NewBuffer(person2JSON))
}
