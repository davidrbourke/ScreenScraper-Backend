package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Foo")
	DoRequest("https://sainsburys.co.uk/")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// DoRequest makes a GET Request to the url
func DoRequest(url string) string {
	resp, err := http.Get(url)

	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	check(err)

	fmt.Println(string(body))

	return string(body)
}
