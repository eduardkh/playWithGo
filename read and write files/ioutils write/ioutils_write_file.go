package main

import "io/ioutil"

func main() {
	err := ioutil.WriteFile("ioutil file.txt", []byte("Hello world"), 0644)
	if err != nil {
		panic("unable to write to file")
	}
}
