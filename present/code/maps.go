package main

import "fmt"

func main() {
	scores := map[string]int{
		"Asha":  88,
		"Budi":  75,
		"Cheng": 92,
	}

	fmt.Println("Asha's score:", scores["Asha"])

	// Check if a key exists
	score, ok := scores["Diana"]
	if !ok {
		fmt.Println("Diana not found, got zero value:", score)
	}
}
