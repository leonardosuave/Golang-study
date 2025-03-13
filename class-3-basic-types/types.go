package main

import (
	"errors"
	"fmt"
)

func main() {
	//	===Type INT===
	// integers
	// every int start with value 0 if not attribute value

	// int8, int16, int32 int64
	// int = computer default. My computer is 64bits, so int = int64
	var number1 int32 = 10000
	number2 := 85555555555
	fmt.Println(number1)
	fmt.Println(number2)
	//=================

	
	//	===Type UINT===
	// every uint start with value 0 if not attribute value

	// uint8, uint16, uint32 uint64 (are numbers without signals)
	// uint = computer default. My computer is 64bits, so uint = uint64
	var number3 uint8 = 100
	number4 := 1000000
	fmt.Println(number3)
	fmt.Println(number4)
	//=================

	
	//	===Type ALIAS===

	//	To int32 and uint8, can declar the type with another name
	//	int32 = rune
	// uint8 = 
	var number5 rune = 855555
	var number6 byte = 85
	fmt.Println(number5)
	fmt.Println(number6)
	//=================

	// ===FLOAT===
	// real numbers
	// every float start with value 0 if not attribute value

	//	float32 or float64
	//	float not exist = computer default. My computer is 64bits, so float64
	var realNumber1 float32 = 1524.589
	realNumber2 := 1855.578
	fmt.Println(realNumber1)
	fmt.Println(realNumber2)
	//=================

	// ===STRING===
	// every string start with value empty if not attribute value

	var str string = "str1"
	str2 := "str2"
	fmt.Println(str)
	fmt.Println(str2)
	//=================

	// ===BOOLEAN===
	// every boolean start with value false if not attribute value

	var isValue1 bool = true
	isValue2 := false
	fmt.Println(isValue1)
	fmt.Println(isValue2)
	//=================

	// ===ERROR===
	// every error start with value nil if not attribute value (nil === null from javascript)
	// To create a error type its neccessary use native package GO, name is errors

	var error1 error = errors.New("Error 1")
	error2 := errors.New("Error 2")
	fmt.Println(error1)
	fmt.Println(error2)
	//=================
}