package main

import "fmt"

func function1() {
	fmt.Println("Executing first function.....")
}

func function2() {
	fmt.Println("Executing second function.....")
}

func main() {
	defer function1() // last thing executed
	function2()
}