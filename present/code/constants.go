package main

import "fmt"

const Pi = 3.14159
const MaxStudents = 40

func main() {
	radius := 5.0
	area := Pi * radius * radius
	fmt.Printf("Area: %.2f\n", area)

	x := 10
	y := float64(x) * Pi // explicit type conversion required
	fmt.Println(y)
}
