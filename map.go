package main

import "fmt"

func main() {
	var month map[int]string = map[int]string{}

	month[1] = "Jan"
	month[2] = "Feb"
	fmt.Println(month)
}
