package main

import "fmt"

func main() {
	defer fmt.Println("step 2")
	fmt.Println("step 1")
}
