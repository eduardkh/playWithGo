package main

import "fmt"

var a int64

func main() {
	a = 64
	fmt.Println("a is =>", a)
	fmt.Printf("a type is => %T\n", a)

	b := 32
	fmt.Println("b is =>", b)
	fmt.Printf("b type is => %T\n", b)

	c := int64(64)
	fmt.Println("c is =>", c)
	fmt.Printf("c type is => %T\n", c)
}
