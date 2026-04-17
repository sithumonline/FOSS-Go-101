package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (r Rectangle) Area() float64   { return r.Width * r.Height }

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	printArea(Circle{Radius: 3})
	printArea(Rectangle{Width: 4, Height: 5})
}
