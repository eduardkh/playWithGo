package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 4},
		{1, 2, 4},
		{1, 2, 4, 1, 2, 48},
	}
	// fmt.Println(a[1][2])
	for b := range a {
		fmt.Println(a[b][len(a[b])-1])
	}
}
