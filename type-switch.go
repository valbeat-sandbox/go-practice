package main

import "fmt"

type Stringer interface {
	String() string
}

// Type Assertion
func TypeAssertion(value interface{}) {
	s, ok := value.(string) // Type Assertion
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Printf("value is not string\n")
	}
}

// Type Switch
func TypeSwitch(value interface{}) {
	switch v := value.(type) {
	case string:
		// typeの指定
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("value is int: %s\n", v)
	case Stringer:
		// interfaceの指定
		fmt.Printf("value is Striger: %s\n", v)
	}
}

func main() {
	TypeAssertion("a")
	TypeAssertion(1)
	TypeSwitch("a")
}
