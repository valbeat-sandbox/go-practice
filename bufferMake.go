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
	// データが読み出されるまでブロック
	ch<- "foo"
}
