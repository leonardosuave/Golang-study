package main

import "fmt"

func changeSignal(n int) int {
	return n * -1
}

func changePointerSignal(n *int) { //Here will change the memory address value. Not neccessary return value
	*n = *n * -1
}

func main() {
	n := 15
	value1 := changeSignal(n) // Create a COPY and not a reference as pointer
	fmt.Println(n, value1)
	fmt.Println()
	nPointer := 20
	changePointerSignal(&nPointer) //&nPointer is the memory address passed by reference 
	fmt.Println(nPointer)
}