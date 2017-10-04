package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://google.com",
		"http://yahoo.com",
		"http://example.com",
	}
	statusChan := make(chan string)
	for _,url := range urls  {
		// waitGroupにカウント追加
		wait.Add(1)
		// 処理を関数化し、goを付けると非同期処理となる
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			// チャネルに書き込み
			statusChan <- res.Status
			// カウント減らす
			wait.Done()
		}(url)
	}
	for i := 0; i < len(urls); i++  {
		// チャネルから読み出し
		fmt.Println(<-statusChan)
	}
	// 待ち合わせ
	wait.Wait()
}
