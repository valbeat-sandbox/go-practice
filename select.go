package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
)


// ステータスをチャネルとして取得するメソッド
// 戻り値を<-chanと読み出し専用にすることで外部からの書き込みを防ぐ
func getStatus(urls []string) <-chan string {
	statusChan := make(chan string)
	for _,url := range urls  {
		// 処理を関数化し、goを付けると非同期処理となる
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			// チャネルに書き込み
			statusChan <- res.Status
		}(url)
	}
	return statusChan
}

func main() {
	timeout := time.After(time.Second)
	urls := []string{
		"http://google.com",
		"http://yahoo.com",
	}
	statusChan := getStatus(urls)
	// 任意のラベル
	LOOP:
	for {
		select {
		case status := <- statusChan:
			fmt.Println(status)
		case <-timeout:
			// for文から抜けるためにラベルを付ける
			break LOOP
		}
	}
}
