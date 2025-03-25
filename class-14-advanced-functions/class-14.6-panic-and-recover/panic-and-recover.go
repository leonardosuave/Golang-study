package main

import "fmt"

func recoverExecution() {

	// create variable and try recover. If variable is different than nil(null), so the things were recovered
	if r := recover(); r != nil {
		fmt.Println("Execution recovered")
	}
}

func approveStudant(n1, n2 float64) bool {
	defer recoverExecution()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media is 6!")
}

func main() {
	fmt.Println(approveStudant(6, 6))
	fmt.Println("After execution")
}