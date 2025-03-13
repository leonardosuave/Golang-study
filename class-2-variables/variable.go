package main

import "fmt"

func main() {
	var variable1 string = "Test variable by var"
	variable2 := "Test variable without var"

	var (
		variable3 string = "variable3"
		variable4 string = "variable4"
	)

	variable5, variable6 := "variable5", "variable6"

	const constant1 string = "constant1"

	//In GO can be change values between 2 variables without a new variable to save value.
	//Ex: variable5 received variable6 value and variable6 received variable5 value 
	// variable5, variable6 = variable6, variable5

	fmt.Println(variable1)
	fmt.Println(variable2)
	fmt.Println(variable3, variable4)
	fmt.Println(variable5, variable6)
	fmt.Println(constant1)
}