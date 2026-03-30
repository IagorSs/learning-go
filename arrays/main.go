package main

import "fmt"

// Arrays with fixed size, will throw panic if try put element out of bounds
var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10
	// Initial values to int is 0
	fmt.Println(x)
	fmt.Println(y)

	// Same type with different size are incompatible types
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// Operations with arrays
	fmt.Println(len(x))
}
