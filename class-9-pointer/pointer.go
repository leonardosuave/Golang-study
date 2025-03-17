package main

import "fmt"

func main() {

	// COPY
	var variable1 int = 10
	var variable2 int = variable1	// Create a copy and not reference by memory address

	fmt.Println(variable1, variable2)

	variable1 = 15

	// variable1 value different variable2 after edit
	fmt.Println(variable1, variable2)
	// =============

	// POINTER
	var variable3 int = 100
	var variable4 *int = &variable3 	// *int default is nil

	fmt.Println(variable3, variable4) 	// Will show variable3 value and to variable4 will show the memory address
	fmt.Println(variable3, *variable4)	// Will show variable3 and variable4 values - *variable4 allow show the memory address value
}