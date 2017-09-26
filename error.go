package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	n, err := div(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("divied by zero")
	}
	return i / j, nil
}
