package main

import "fmt"

func main() {
	x := [5]int{0, 1, 2, 3, 4}

	for _, value := range x {
		fmt.Println(value)
	}

	fmt.Printf("%T", x)
}
