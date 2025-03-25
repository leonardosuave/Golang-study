package main

import "fmt"

// This way allow received a lot of int params, like spreed operator from JS - allowed received none params
func sum(n ...int) int {
	fmt.Println(n)
	total := 0
	for _, number := range n {
		total += number
	}

	return total
}

func concatSrtingNumber(txt string, n ...int) {
	for _, number := range n {
		fmt.Println(txt, number)
	}
}

func main() {
	totalSum := sum(1, 14, 12, 43, 2, 3, 1)
	fmt.Println(totalSum)
	concatSrtingNumber("Some here", 1, 1, 2, 10, 5)
}