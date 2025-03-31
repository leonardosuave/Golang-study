package main

import "fmt"

type generic interface {
}

// interfaces empty are generic, like any js
func genericFunction(params generic) {
	fmt.Println(params)
}

func main() {
	booleanVariable := false
	genericFunction(booleanVariable)

	stringVariable := "Leonardo"
	genericFunction(stringVariable)

	numberVariable := 99999
	genericFunction(numberVariable)

}