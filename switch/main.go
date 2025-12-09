package main

import "fmt"

func main() {
	value := 3

	switch value {
	case 1:
		fmt.Println("Value is one")
	case 2:
		fmt.Println("Value is two")
	case 3:
		fmt.Println("Value is three")
	}

	switch {
	case value > 2:
		fmt.Println("Value is greater than two")
		fallthrough
	// Will execute this because of fallthrough, don't check condition
	case value > 5:
		fmt.Println("Value is greater than five")
	// Don't execute this because already matched above
	case value == 3:
		fmt.Println("Value is three")
	default:
		fmt.Println("Value is something else")
	}
}
