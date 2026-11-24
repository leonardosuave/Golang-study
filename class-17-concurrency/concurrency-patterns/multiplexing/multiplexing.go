package main

import "fmt"

func main() {
	channel := multiplexing(write("Hello world 1"), write("Hello world 2"))

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func multiplexing(channelIn1, channelIn2 <-chan string) <-chan string {
	// channel to receive data from multiple sources
	channelOut := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelIn1:
				channelOut <- message
			case message := <-channelIn2:
				channelOut <- message
			}
		}
	}()

	return channelOut	
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