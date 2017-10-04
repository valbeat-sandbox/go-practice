package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://yahoo.com",
		"http://example.com",
	}
	for _,url := range urls  {
		// 処理を関数化し、goを付けると非同期処理となる
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
		}(url)
	}
	// main()が終わらないようにSleep
	time.Sleep(time.Second * 5)
}
