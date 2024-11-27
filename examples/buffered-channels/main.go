package main

import "fmt"

func main() {
	messages := []string{"first msg", "second msg", "third msg"}

	// this channel only has capacity for 3
	// but we will use concurrency to share memory and send the latest message
	// we can use buffered channels so we never have to wait for the receiver to store the data
	// we store the data in the buffer, if the buffer is full we block
	bufferedChannel := make(chan string, 3)

	messages = addMessageToSlice("last msg", messages)
	// by enabling concurrency by using "go" keyword in front of this function
	// we send the latest message to the channel when the first one is out
	// if we remove the "go" keyword we will have a deadlock!
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
