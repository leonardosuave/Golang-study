package main

import (
	"fmt"
	"reflect"
)

func main() {

	//ARRAYS
	// The arrays has fixed positions and types, so i need to declair how many positions the array will have and all index will have the unic type declaired 

	var array1 [5] string
	array1[1] = "Some thing here"
	fmt.Println(array1)

	array2 := [5]string{"some1", "some2", "some3", "some4", "some5"}
	fmt.Println(array2)

	// [...] To array have index quantity equal the values decalired
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	//========

	//SLICES
	// Slices are flexibler than arrays, but they are not a array. (Its allow only one type to all slice and slices not have size limit)
	slice1 := []int{10, 11, 12}
	fmt.Println(slice1)

		// To see that arrays and slices are not equals
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	slice1 = append(slice1, 15)	// as a push() js
	fmt.Println(slice1)

	slice2 := array2[1:3]	// To take values from a array
	fmt.Println(slice2)

	// Intern Arrays (make func - reserve memory space to some thing)
	//OBS make(type, lenght, cap) - cap is optional and to this cap is leght value
	slice3 := make([]float32, 10, 11) // make func create a slice with 11 index and return the 10 first index (cap is optional)
	slice3 = append(slice3, 12)
	slice3 = append(slice3, 15)	// Pass the cap limit 
	fmt.Println(slice3)
	fmt.Println(len(slice3))	
	fmt.Println(cap(slice3))	// When pass the limit, the slice take the size (12) and doble the cap to size 24
}