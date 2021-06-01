package main

import "fmt"

func main() {
	suka := []int{1, 2, 3, 4, 5}
	blat := []int{6, 7, 8, 9, 10, 11, 12, 13, 14}
	suka = append(suka, blat...)
	for _, s := range suka {
		fmt.Printf("%d \n", s)
	}
}
