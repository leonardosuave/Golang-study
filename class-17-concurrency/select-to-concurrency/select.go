package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	// It will have problem, because channel1 can print 4 times in 2 seconds, but channel2 will print only 1 time in 2 seconds, so the program will wait for channel2 to print and then print channel1 again. So need to use selec to avoid this problem.

	// go func() {
	// 	for {
	// 		time.Sleep(time.Millisecond * 500)
	// 		channel1 <- "Channel 1"
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 2)
	// 		channel2 <- "Channel 2"
	// 	}
	// }()

	// for {
	// 	message1 := <-channel1
	// 	fmt.Println(message1)

	// 	message2 := <-channel2
	// 	fmt.Println(message2)
	// }

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		// Now the channel are independent and run separately.
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		}
	}
}
