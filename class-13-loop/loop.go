package main

import (
	"fmt"
	"time"
)

func main() {
	increment := 0
	for increment < 10 {
		time.Sleep(time.Second) // Its a 1 sec timeout, like setTimeOut JS 
		fmt.Println(increment)
		increment++
	}

	for increment1 := 0; increment1 < 5; increment1++ {
		time.Sleep(time.Second)
		fmt.Println(increment1)
	}
	
	increment2 := 0
	sliceValues := []string{"some 1", "some 2", "some 3", "some 4", "some 5"}
	for increment2 < len(sliceValues) {
		time.Sleep(time.Second)
		fmt.Println(sliceValues[increment2])
		increment2++
	}

	// FOR by RANGE (as for of oe forEach JS)

	names := [3]string{"leo", "tassi", "judite"}
	for index, value := range names {	// need to pass 2 attribute (first = index and second = value). To pass only value, declar index as _
		time.Sleep(time.Second)
		fmt.Println(index, value)
	}

	for index, value := range "WORD" {
		fmt.Println(index, string(value)) // If not pass string(value), will show the ASC table number
	}

	user := map[string]string {
		"name": "Leonardo",
		"nick": "Leo",
	}
	for key, value := range user {
		fmt.Println(key, value)
	}

}