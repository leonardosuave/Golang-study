package main

import "fmt"

func main() {
	func(txt string) {
		fmt.Println(txt)
	}("params")

	r := func(txt string, number int) string {
		return fmt.Sprintf("Receiving -> %s %d", txt, number) //To concat values. Need to declair the types concated by % -> %s is a string and %d is a int
	}("n params", 10)
	fmt.Println(r)
}