package main

import "fmt"

func main() {
	_string := "this is a string"
	_stringPointer := &_string
	fmt.Println(_string)
	fmt.Printf("Pointer: %v\n", _stringPointer)
	fmt.Printf("Dereference (get the value): %v", *_stringPointer)
}
