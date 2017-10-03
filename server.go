package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello, world")
}

func main() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
