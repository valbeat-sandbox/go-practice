package main

import (
	"fmt"
	"log"
	"net/http"
)

// ステータスをチャネルとして取得するメソッド
// 戻り値を<-chanと読み出し専用にすることで外部からの書き込みを防ぐ
func getStatus(urls []string) <-chan string {
	// バッファをURLの数(3)に
	statusChan := make(chan string,3)
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
	urls := []string{
		"http://google.com",
		"http://yahoo.com",
		"http://example.com",
	}
	statusChan := getStatus(urls)
	for i := 0; i < len(urls); i++  {
		// チャネルから読み出し
		fmt.Println(<-statusChan)
	}
}
