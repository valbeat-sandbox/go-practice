package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./hoge.go")
	if err != nil {
		// エラー処理
		fmt.Println(err)
		return
	}
	// 関数抜ける直前に絶対実行される
	defer file.Close()
	// 正常処理
}
