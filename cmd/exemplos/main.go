package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { // goroutine 1
	ch := make(chan int) // Create a channel
	qtdWorkers := 3

	// //goroutine 2
	// go func() {
	// 	ch <- "Hello, World!" // Send a message to the channel
	// }()

	// msg := <-ch // Receive the message from the channel
	// fmt.Println(msg)

	for i := range qtdWorkers {
		go worker(i, ch)
	}

	for i := range 10 {
		ch <- i
	}
}
