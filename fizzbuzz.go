package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(ToFizzBuzz(i))
	}
}
func ToFizzBuzz(amount int) interface{} {
	if amount%15 == 0 {
		return "FizzBuzz"
	} else if amount%3 == 0 {
		return "Fizz"
	} else if amount%5 == 0 {
		return "Buzz"
	} else {
		return amount
	}
}
