package main

import "fmt"

type Student struct {
	Name  string
	Score float64
}

func main() {
	// Array — fixed size, rarely used directly.
	var topThree [3]string
	topThree[0] = "Asha"
	topThree[1] = "Budi"
	topThree[2] = "Cheng"

	// Slice of structs — dynamic, resizable, the workhorse of real Go code.
	students := []Student{
		{"Asha", 88.5},
		{"Budi", 75.0},
	}

	students = append(students, Student{"Cheng", 92.0})

	fmt.Println(topThree)
	fmt.Println(students)
}
