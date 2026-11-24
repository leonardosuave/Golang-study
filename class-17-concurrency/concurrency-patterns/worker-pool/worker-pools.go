package main

import "fmt"

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results) //go routine (dont need finish to run next line code)

	// Now the program will run 2 go routines, one for each worker, and will run concurrently, and will have better performance, but is limited by the number of cores of the machine.
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i // send the task to tasks channel
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results // receive the result from results channel
		fmt.Println(result)
	}
}

// tasks receive data only and results send data only
func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number) // send the result to results channel
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}