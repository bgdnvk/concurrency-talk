package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	// wait 2s and send data to the channel
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Data from Goroutine"
	}()

	// infinite loop
	for {
		select {
		case msg := <-channel:
			fmt.Println(msg)
			return // break infinite loop

		// default case simply simulates some work
		// we call it every time  until we get the other case
		default:
			fmt.Println("No data yet, doing other work...")
			time.Sleep(500 * time.Millisecond) // Simulate some work
		}
	}
}
