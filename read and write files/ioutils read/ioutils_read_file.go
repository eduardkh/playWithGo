package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	buf, err := ioutil.ReadFile("ioutil file.txt")
	if err != nil {
		panic("unable to read the file")
	}
	fmt.Println(string(buf))
}
