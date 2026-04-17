package main

import "fmt"

// This does NOT modify the original
func doubleByValue(n int) {
	n *= 2
}

// This DOES modify the original
func doubleByPointer(n *int) {
	*n *= 2
}

func main() {
	x := 10
	doubleByValue(x)
	fmt.Println("by value:", x) // still 10

	doubleByPointer(&x)
	fmt.Println("by pointer:", x) // now 20
}
