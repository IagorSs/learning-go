package main

import "fmt"

// Custom type based on int
type hotdog int

func main() {
	var x hotdog = 55
	y := 55

	fmt.Println(x)
	fmt.Println(y)

	// Type conversion (x=y doesn't work)
	x = hotdog(y) + 88
	fmt.Println(x)
}
