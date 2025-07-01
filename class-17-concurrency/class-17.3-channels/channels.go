package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("Hello world", channel)


	for {
		message, statusChannel := <-channel // channel receive value (<-) and save in the const message (when specify that will receive a value, its the same to say that need to wait receive to continue the code(will sync the goroutine))

		// check if channel was closed avoid deadlocks, as is in a infinity looping and write is a for only 5 times (in the 6 looping, message will not receive value and will generate a deadlock) 
		if !statusChannel {
			break
		}
		fmt.Println(message) // print the channel value received

		
		// **BEST SYNTAX to check if channel is receiving value and avoid deadlocks
		/*
			for message := range channel {
				fmt.Println(message)
			}
		*/
	}

	fmt.Println("end programming")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {

		channel <- text // send text to channel 
		time.Sleep(time.Second)
	}

	close(channel) // close the channel and avoid deadlock
}