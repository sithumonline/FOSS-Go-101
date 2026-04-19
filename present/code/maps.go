package main

import "fmt"

type Student struct {
	Name  string
	Score float64
}

func main() {
	// Map from student name to their record.
	roster := map[string]Student{
		"asha":  {"Asha", 88.5},
		"budi":  {"Budi", 75.0},
		"cheng": {"Cheng", 92.0},
	}

	fmt.Println("Asha's score:", roster["asha"].Score)

	// Two-value form: check whether a key exists.
	if s, ok := roster["diana"]; ok {
		fmt.Println("found:", s)
	} else {
		fmt.Println("Diana is not on the roster yet")
	}
}
