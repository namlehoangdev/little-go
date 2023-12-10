package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to receive data
	dataCh := make(chan string)

	// Create a timeout channel
	timeoutCh := time.After(2 * time.Second)

	// Launch a goroutine to send data after a delay
	go func() {
		time.Sleep(100 * time.Second)
		dataCh <- "Data received!"
	}()

	// Use select to wait for data or timeout
	select {
	case data := <-dataCh:
		fmt.Println("Received data:", data)
	case t := <-timeoutCh:
		fmt.Println("Timed out waiting for data", t)
	}
}
