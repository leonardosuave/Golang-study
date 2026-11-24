package main

import "fmt"

func main() {
	channel := write("Hello world") // receive the channel from the function write
	fmt.Println(<-channel)
}

// Return a channel
func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("value received: %s", text) // Attribute the value to the channel, and the channel will be blocked until the value is received
		}
	}()

	return channel
}