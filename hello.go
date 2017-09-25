package main

import "fmt"

func main() {
	a := 100
	b := 11.3
	c := "hoge"
	var d bool
	// データ型に応じたフォーマットを選択する
	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)
}
