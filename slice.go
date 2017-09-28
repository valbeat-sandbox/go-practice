package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c"}
	fmt.Println(slice[0])
	slice = append(slice, "d")
	fmt.Println(slice[3])

	for i, s := range slice {
		fmt.Println(i, s)
	}
}
