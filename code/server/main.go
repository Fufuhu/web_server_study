package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")

	fmt.Printf("%v\n", r)

	for field, values := range r.Header {
		fmt.Printf("%s\n", field)
		for _, value := range values {
			fmt.Printf("	Value=%s\n", value)
		}
	}

	body_closer := r.Body
	bytes_buffer := new(bytes.Buffer)

	bytes_buffer.ReadFrom(body_closer)
	fmt.Printf("%v\n", bytes_buffer.String())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
