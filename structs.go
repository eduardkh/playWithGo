package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

// greeting method for person struct (value receiver)
func (p person) greet() string {
	return fmt.Sprintf("Hello my name is %s %s and i am %d years old", p.FirstName, p.LastName, p.Age)
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
	fmt.Println(per.greet())
}
