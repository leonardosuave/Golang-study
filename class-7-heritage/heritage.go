package main


import "fmt"

type person struct {
	name string
	age uint8
	height float32
}

type studant struct {
	person
	curse string
	university string
}

func main() {
	fmt.Println("heritage")
	person1 := person{"Leonardo", 29, 1.78}
	fmt.Println(person1)

	studant1 := studant{person1, "ADS", "UFSCar"}
	fmt.Println(studant1)
	
	fmt.Println(studant1.name)
}