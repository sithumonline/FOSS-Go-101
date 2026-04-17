package main

import "fmt"

func main() {
	// Array — fixed size
	var arr [3]string
	arr[0] = "Go"
	arr[1] = "is"
	arr[2] = "fun"

	// Slice — dynamic, used in real Go code
	marks := []int{85, 90, 78}
	marks = append(marks, 95)

	fmt.Println(arr)
	fmt.Println(marks)
}
