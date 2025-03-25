package main

import "fmt"

func closure() func() {
	txt := "Inside closure function"

	function := func() {
		fmt.Println(txt)
	}

	return function
}

func main() {
	// Are function that reference variable outside body(scope)
	txt := "Inside main function"
	fmt.Println(txt)

	newFunction := closure() //newFunction variable becames function type, so i can execute like a function
	newFunction()

}