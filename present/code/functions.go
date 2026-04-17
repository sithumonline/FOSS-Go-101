package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println(add(3, 4))

	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%.2f\n", result)
}
