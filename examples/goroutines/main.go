package main

import (
	"fmt"
	"time"
)

// print the ints in the loop
func printFrom(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	printFrom("direct")

	// if you keep launching the program sometimes you will see goroutine2 first
	go printFrom("goroutine1: runs concurrently")
	go printFrom("goroutine2: runs concurrently")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous goroutine")

	// wait a second for goroutines to complete
	// otherwise the program finishes without displaying results from our goroutines
	time.Sleep(time.Second)
	fmt.Println("done")
}
