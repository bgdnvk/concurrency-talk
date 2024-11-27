package main

import (
	"fmt"
)

func main() {
	messages := []string{"first msg", "second msg", "third msg"}
	moreMessages := []string{"fourth msg", "fifth msg", "last msg"}

	bufferedChannel := make(chan string, 3)

	// due to the concurrent nature of our program we can just keep sending messages to the channel
	for _, m := range moreMessages {
		messages = addMessageToSlice(m, messages)
	}

	go sendMessagesToChannel(messages, bufferedChannel)

	for m := range bufferedChannel {
		fmt.Println(m)
	}
}

func sendMessagesToChannel(messages []string, bufchan chan string) {
	for _, m := range messages {
		bufchan <- m
	}
	close(bufchan)
}

func addMessageToSlice(msg string, messages []string) []string {
	return append(messages, msg)
}
