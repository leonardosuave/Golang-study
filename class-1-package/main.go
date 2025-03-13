package main

import (
	"github.com/badoux/checkmail"
	"module/helper-module"
	"fmt"
)

func main() {
	fmt.Println("First code with module")
	helpermodule.Write()
	// helpermodule.writeTwo() - Not allowed because its a first lowercase word and from another package
	erro := checkmail.ValidateFormat("leonardo@email.com")
	fmt.Println(erro)
}