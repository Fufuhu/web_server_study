package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.google.co.jp"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
