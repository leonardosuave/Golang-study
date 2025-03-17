package main

import "fmt"

func main() {
	number1 := 16

	if number1 < 10 {
		fmt.Println("Smaller than 10")
	} else {
		fmt.Println("Higher than 10")
	}

	// otherNumber is create to use only into ifs. Outside otherNumber not exist
	if otherNumber := number1; (otherNumber > 10 && otherNumber <= 15){
		fmt.Println("Othernumber between 10 and 15")
	} else if (otherNumber > 15) {
		fmt.Println("Othernumber higher than 15")
	} else {
		fmt.Println("Othernumber smaller or equal 10")
	}
}