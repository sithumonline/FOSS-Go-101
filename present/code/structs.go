package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

func main() {
	s := Student{Name: "Asha", Age: 20, Score: 88.5}
	fmt.Println(s.Name, "scored", s.Score)

	students := []Student{
		{Name: "Budi", Age: 21, Score: 75.0},
		{Name: "Cheng", Age: 19, Score: 92.0},
	}
	for _, st := range students {
		fmt.Printf("%s: %.1f\n", st.Name, st.Score)
	}
}
