package main

import "fmt"

func main() {
	// Aritimetic

	// The number need to be same type. Ex int16 = 12 and int8 = 10 can not be used together to operators because are diferent types 
	sum := 1 + 2
	subtraction := 2 - 1
	division := 10 / 5
	divisionRemains := 10 % 2

	fmt.Println(sum, subtraction, division, divisionRemains)
	//============

	//	Attribute
	var variable1 string = "variable1"
	const constante1 string = "constante"
	variable2 := "variable2"
	constante2 := "constante2"

	fmt.Println()
	println(variable1, constante1, variable2, constante2)
	//============

	// Relation operators (true or false)
	fmt.Println()
	fmt.Println(1 > 2)	// 	higher
	fmt.Println(1 >= 2)	//	higher or equal
	fmt.Println(1 < 2)	//	smaller
	fmt.Println(1 <= 2)	// 	smaller or equal
	fmt.Println(1 == 2)	// 	eqaul
	fmt.Println(1 != 2) //	Different
	//============

	// Logics operators
	// Work equal javascript
	fmt.Println()
	variableTrue, variableFalse := true, false
	fmt.Println(variableTrue && variableFalse)
	fmt.Println(variableTrue || variableFalse)
	fmt.Println(!variableTrue)
	fmt.Println(!variableFalse)
	//============

	// Unario operators
	fmt.Println()
	variableNumber := 10
	variableNumber ++	//	equal js
	fmt.Println(variableNumber)
	variableNumber2 := 15
	variableNumber2 += 10 //	equal js
	fmt.Println(variableNumber2)
	//============

	//	Ternário (not exist in GO)
	// 	To do the same, need to be by if/else
	var stringValue string
	var numberValue int16 = 15
	// Work equal js or without ()
	if numberValue > 10 {
		stringValue = "Maior que 10"
	} else {
		stringValue = "Menor que 10"
	}
	fmt.Println(stringValue)
}