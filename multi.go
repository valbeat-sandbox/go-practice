package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
)

var empty struct{} // サイズがゼロの構造体

// ステータスをチャネルとして取得するメソッド
// 戻り値を<-chanと読み出し専用にすることで外部からの書き込みを防ぐ
func getStatus(urls []string) <-chan string {
	// バッファをURLの数(3)に
	statusChan := make(chan string, 3)
	limit := make(chan struct{}, 5)
	go func() {
		for _,url := range urls {
			select {
			case limit <- empty:
				// limitに書き込みが可能な場合は取得処理を実行
				go func(url string) {
					// このゴルーチンは同時に5つのみ起動
					res, err := http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					defer res.Body.Close()
					// チャネルに書き込み
					// mainの読み込みが遅くてもbufferがあるのですぐに終わる
					statusChan <- res.Status
					// 終わったら一つ解放
					<-limit
				}(url)
			}
		}
	}()
	return statusChan
}

func main() {
	timeout := time.After(time.Second * 10)
	urls := []string{
		"http://google.com",
		"http://yahoo.com",
		"http://exampl.com",
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
