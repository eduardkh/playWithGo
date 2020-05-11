package main

import "fmt"

func main() {
	// the array is fixed light slices
	array := [4]int{1, 2, 3, 4}

	// slices can be modified
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = append(slice, 11)
	fmt.Println(array)
	fmt.Println(slice)
}
