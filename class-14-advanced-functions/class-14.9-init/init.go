package main

import "fmt"

// variable global
var n int 

// init function is the first to be executed. Its first than main.
// it can have 1 init function per file. Different than main that can have only 1 per package
func init() {
	fmt.Println("first")
	n = 10
}

func main() {
	fmt.Println("second")
	fmt.Println(n)
}