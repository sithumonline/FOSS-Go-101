package main

import "fmt"

func main() {
	names := []string{"Asha", "Budi", "Cheng"}

	for i, name := range names {
		fmt.Printf("%d: %s\n", i, name)
	}
}
