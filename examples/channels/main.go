package main

import "fmt"

func main() {
	fmt.Println("starting the program")
	// make a new channel
	messages := make(chan string)

	// use an anymous  goroutine to send the string "ping" to the channel
	go func() { messages <- "first msg" }()
	go func() { messages <- "second msg" }()

	// we only loop twice becase we know we sent 2 messages
	for i := 0; i < 2; i++ {
		fmt.Println(<-messages)
	}

	// Deadlock!
	// this loop is waiting forever on the channel
	// because the channel is never closed
	//for range messages {
	//	// get the data from the channel
	//	msg := <-messages
	//	// display the data we got
	//	fmt.Println(msg)
	//}
}
