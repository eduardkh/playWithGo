package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

// greeting method for person struct (value receiver) - when getting a value (getter)
func (p person) greet() string {
	return fmt.Sprintf("Hello my name is %s %s and i am %d years old", p.FirstName, p.LastName, p.Age)
}

// hasBirthday method for person struct (pointer receiver) - when changing a value (setter)
func (p *person) hasBirthday() {
	p.Age++
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
	per.hasBirthday()
	fmt.Println(per.greet())
}
