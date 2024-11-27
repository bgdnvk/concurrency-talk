package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	// msg will arrive after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from Channel 1"
	}()

	// msg will arrive after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Message from Channel 2"
	}()

	// we will execute this block twice and check every
	// if we get a msg back
	// then print whichever is ready
	for i := 0; i < 2; i++ {
		// select block waits for one of the channels to receive a msg
		// it waits for at least one channel to be ready
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
			// we can add a default case for when we don't want to wait for a msg
			// default:
			//	fmt.Println("default")
		}
	}
}
