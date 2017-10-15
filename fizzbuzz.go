package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		ToFizzBuzz(i)
	}
}
func ToFizzBuzz(amount int) {
	if amount%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if amount%3 == 0 {
		fmt.Println("Fizz")
	} else if amount%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(amount)
	}
}
