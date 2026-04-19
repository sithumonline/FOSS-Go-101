package main

import "fmt"

// greet runs in its own goroutine and signals completion on the channel.
func greet(name string, done chan bool) {
	fmt.Println("Hello,", name)
	done <- true
}

func main() {
	done := make(chan bool)

	// Start three goroutines — they run concurrently.
	go greet("Amina", done)
	go greet("Budi", done)
	go greet("Cheng", done)

	// Wait for all three to finish.
	<-done
	<-done
	<-done
}
