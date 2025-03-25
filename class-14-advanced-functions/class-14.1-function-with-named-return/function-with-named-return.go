package main

import "fmt"

// The variables are created in type function and not neccessary return variable because its in the type function
func calculate(n1, n2 int) (sum, subtraction int) {
	sum = n1 +n2
	subtraction = n1 - n2
	return
}

func main() {
	sum1, subtraction1 := calculate(10, 15)
	fmt.Println(sum1, subtraction1)
}