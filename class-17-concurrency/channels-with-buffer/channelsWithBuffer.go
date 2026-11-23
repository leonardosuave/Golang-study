package main

import "fmt"

func main() {
	// This case only blocked with deadlock when the buffer is full, because the channel is buffered and can hold 2 values before blocking
	channel := make(chan string, 2) // create a channel with buffer of 2

	channel <- "Hello world"
	channel <- "Hello world 2"

	message := <-channel
	message2 := <-channel

	fmt.Println(message)
	fmt.Println(message2)
}