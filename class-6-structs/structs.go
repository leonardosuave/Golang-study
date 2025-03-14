package main

import "fmt"

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	street string
	n uint16
}

func main() {
	fmt.Println("Structs file")

	var user1 user
	fmt.Println(user1)

	user1.name = "Leonardo"
	user1.age = 29
	fmt.Println(user1)

	// By inference
	// user2 := user{"Suave", 29, address{"Terezina", 1120}}
	address := address{"Terezina", 1120}
	user2 := user{"Suave", 29, address}
	fmt.Println(user2)

	// To optional value
	user3 := user{age: 29}
	fmt.Println(user3)
}