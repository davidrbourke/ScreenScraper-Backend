package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Foo")

	resp, err := http.Get("https://sainsburys.co.uk/")

	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	check(err)

	fmt.Println(string(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
