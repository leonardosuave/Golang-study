package main

import (
	"commander-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("start!")

	aplication := app.Generate() // Imported from another file with the logic to run.
	if error := aplication.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}