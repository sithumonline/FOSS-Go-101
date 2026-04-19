package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// Multiple return values are common in Go.
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println("3 + 4 =", add(3, 4))

	q, r := divmod(17, 5)
	fmt.Println("17 / 5 =", q, "remainder", r)
}
