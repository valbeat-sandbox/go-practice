package main

import "fmt"

func main() {
	a := 100
	b := 11.3
	c := "hoge"
	var d bool
	// データ型に応じたフォーマットを選択する
	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)

	// 定数はconst
	const name = "takuma"
	fmt.Println(name)

	// iota = enum ?
	const (
		sun = iota // 0
		mon
		tue
	)
	fmt.Println(sun, mon, tue)

	a1, b1 := 10, 100
	if a1 > b1 {
		fmt.Println("a > b")
	} else if a1 < b1 {
		fmt.Println("a < b")
	} else {
		fmt.Println("a = b")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("無限ループ")
	}
}
