package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		// 一秒後にデータを書き込む
		ch <- "foo"
	}()
	// データが書き込まれるまでブロック
	fmt.Println(<-ch)
}
