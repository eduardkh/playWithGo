package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("files.txt")
	if err != nil {
		panic("unable to create file")
	}
	defer f.Close()
	cnt, err := f.WriteString("Hello world")
	if err != nil {
		panic("unable to write to file")
	}
	fmt.Printf("wrote %d bytes\n", cnt)
}
