package main

import (
	"fmt"
	"sync"
)

func main() {
	messages := []string{"first msg", "second msg", "third msg"}
	moreMessages := []string{"fourth msg", "fifth msg", "last msg"}

	bufferedChannel := make(chan string, 3)

	// we make an object to control the flow of our execution
	// it contains a counter for each executed function
	wg := sync.WaitGroup{}
	for _, m := range moreMessages {
		// every time we spin a new goroutine we increase the counter
		// on how much we have to wait
		wg.Add(1)
		// keep in mind the messages will not arrive in order because go routines spawn concurrently (non-deterministic behaviour)
		go func(m string) {
			messages = addMessageToSlice(m, messages)
			// once the execution is done we let our WaitGroup know (this reduces the counter)
			wg.Done()
		}(m)

	}

	// we make sure we wait for the previous block to finish
	wg.Wait()

	go sendMessagesToChannel(messages, bufferedChannel)

	for m := range bufferedChannel {
		fmt.Println(m)
	}
}

func sendMessagesToChannel(messages []string, bufchan chan string) {
	for _, m := range messages {
		bufchan <- m
	}
	// remember to always close the channel
	close(bufchan)
}

func addMessageToSlice(msg string, messages []string) []string {
	messages = append(messages, msg)
	return messages
}
