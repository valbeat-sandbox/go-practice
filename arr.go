package main

import "fmt"

func printArr(arr [4]string) {
	fmt.Println(arr)
}

func main() {
	var arr4 [4]string
	var arr5 [5]string
	printArr(arr4)
	printArr(arr5)
}
