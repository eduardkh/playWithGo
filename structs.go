package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	per := person{
		FirstName: "Eddie",
		LastName:  "Khiaev",
		Age:       36,
	}
	fmt.Println(per)
	fmt.Println(per.FirstName)
	fmt.Println(per.LastName)
	fmt.Println(per.Age)
}
