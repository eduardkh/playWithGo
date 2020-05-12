package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("files.txt")
	if err != nil {
		panic("unable to open file")
	}
	defer f.Close()
	// can reed only 64 chars from file
	buf := make([]byte, 64)
	cnt, err := f.Read(buf)
	if err != nil {
		panic("unable to read from file")
	}
	fmt.Printf("read %d bytes\n", cnt)
	// prints all of the 64 bytes
	// fmt.Println(buf)
	// prints only the bytes we need
	// fmt.Println(buf[:cnt])

	fmt.Println(string(buf[:cnt]))
}
