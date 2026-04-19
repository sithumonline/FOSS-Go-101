package main

import "fmt"

type Student struct {
	Name  string
	Score float64
}

func main() {
	students := []Student{
		{"Asha", 88.5},
		{"Budi", 75.0},
		{"Cheng", 92.0},
	}

	for i, s := range students {
		fmt.Printf("%d. %s — %.1f\n", i+1, s.Name, s.Score)
	}
}
