package main

import "fmt"

var sum func(i, j int) int = func(i, j int) int {
	return i + j
}

func main() {
	fmt.Println(sum(2, 4))
}
