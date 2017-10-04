package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://yahoo.com",
		"http://example.com",
	}
	for _,url := range urls  {
		// http.Getは同期処理のため、レスポンスが返ってくるまで進まない
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		fmt.Println(url, res.Status)
	}
}
