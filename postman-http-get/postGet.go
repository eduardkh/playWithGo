package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://httpbin.org/uuid"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
