package main

import "fmt"

// Only one return
func sum(number1 int8, number2 int8) int8 {
	return number1 + number2
}
// Only one return
func subtractionValues(number1 int8, number2 int8) int8 {
	return number1 - number2
}

// To return two or more values (each value need to be instanced by a variable)
func calculate(number1, number2 int8) (int8, int8) {
	sum := sum(number1, number2)
	subtraction := subtractionValues(number1, number2)

	return sum, subtraction
}

func main() {
	sum1 := sum(15, 18)
	fmt.Println(sum1)

	var f = func(n1 int8, n2 int8) {
		sum2 := sum(n1, n2)
		fmt.Println(sum2)
	}
	f(3, 3)

	//To take only one result, ignored the return order by a _ . Ex: resultSum, _ :=..... to use only the sum result.
	resultSum, resultSubtraction := calculate(12, 10)
	fmt.Println(resultSum)
	fmt.Println(resultSubtraction)
}