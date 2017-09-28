package main

import "fmt"

func main() {
	func(i, j int) {
		fmt.Println(i, j)
	}(2, 4)
}
