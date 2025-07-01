package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WaitGroup: Sincronnyze the goroutines to end before the program end (code end).

	var waitGroup sync.WaitGroup	// instace WaitGroup
	waitGroup.Add(2)	// refere goroutines count in the group that need to wait ends (put goroutines in the queue)

	go func() {
		write("Hello world")
		waitGroup.Done() // Remove a count in the waitGroup queue (total 2 - 1 done)
	}()

	go func() {
		write("Programing in GO lang!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // To wait waitGroup be 0
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}