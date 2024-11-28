package main

import "fmt"

func main() {
	fmt.Println("starting the program")
	// make a new unbuffered channel
	messages := make(chan string)

	// we pass the channel in the function
	// share memory by communicating
	deadlockExample(messages)
	// normalLoop(messages)
}

func deadlockExample(messages chan string) {
	// use an anymous  goroutine to send the string to the channel
	go func() { messages <- "first msg" }()
	go func() { messages <- "second msg" }()
	// Deadlock!
	// this loop is waiting forever on the channel
	// because the channel is never closed
	for msg := range messages {
		// display the data we got
		fmt.Println(msg)
	}
}

func normalLoop(messages chan string) {
	go func() {
		// send the data
		messages <- "first msg"
		messages <- "second msg"
		// close the channel
		close(messages)
	}()

	// now that the channel is closed this for loop won't wait forever
	for msg := range messages {
		// get the data from the channel
		// display the data we got
		fmt.Println(msg)
	}
}
