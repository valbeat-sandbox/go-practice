package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		// 一秒後にデータが読まれる
		<-ch
	}()
	// ブロックしない
	ch<- "1"
	ch<- "2"
	ch<- "3"
	// データが読み出されるまでブロック
	ch<- "4"
	fmt.Println(<-ch)
}
