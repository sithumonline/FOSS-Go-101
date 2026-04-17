package main

import "fmt"

func main() {
	// C-style
	for i := 0; i < 3; i++ {
		fmt.Println("count:", i)
	}

	// while-style
	n := 1
	for n < 8 {
		n *= 2
	}
	fmt.Println("n:", n)
}
