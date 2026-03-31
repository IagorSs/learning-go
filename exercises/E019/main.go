package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 55, 6, 70, 8, 9}

	for _, value := range x {
		fmt.Println(value)
	}

	fmt.Printf("%T %d %d", x, len(x), cap(x))
}
