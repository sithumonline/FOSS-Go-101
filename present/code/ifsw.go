package main

import "fmt"

func grade(score int) string {
	switch {
	case score >= 80:
		return "A"
	case score >= 65:
		return "B"
	case score >= 50:
		return "C"
	default:
		return "F"
	}
}

func main() {
	scores := []int{90, 70, 55, 40}
	for _, s := range scores {
		fmt.Printf("%d → %s\n", s, grade(s))
	}
}
