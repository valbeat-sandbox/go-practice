package main

import "fmt"

func main() {
	var i int = 10
	callByValue(i)
	fmt.Println(i) // 10
	callByRef(&i)  // cと同じように&でアドレスを取得
	fmt.Println(i) // 20
}

func callByValue(i int) {
	i = 20
}

// 型の前に*をつける
func callByRef(i *int) {
	*i = 20
}
