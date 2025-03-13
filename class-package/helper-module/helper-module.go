package helpermodule

import (
	"fmt"
)

// Exported functions need to be comment beforer
func Write() {
	fmt.Println("From helper-module")
	writeTwo() //Allowed only same package because its a first lowercase word
}