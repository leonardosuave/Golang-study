package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCURRENCY != PARALLELISM

	// PARALLELISM: Two or more task executed same time when execute with CPU with one more CORE and the CORES manage the process tasks
	// CONCURRENCY: Two or more task runned same time without wait another task (only 1 CORE in the CPU) (task 1 run a little and stop after another task run a little and stop.....)

	go write("Hello World!") //goroutine
	write("Programing in GO lang!")
}

func write(text string) {
	// For infinity to example
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}