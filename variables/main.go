package main

import "fmt"

// GO is static type language

var y string = "Hello, Go!"

func main() {
	// Explicit declaration
	var x int = 10

	// Short declaration, infers type
	x_f := 10.0
	z := false

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x_f: %v, %T\n", x_f, x_f)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)

	// Need to use this operator, cannot use declaration operator
	x = 20
	fmt.Printf("Updated x: %v, %T\n", x, x)

	// Each type have your zero-value when is declared and not initialized
	var zero int
	fmt.Printf("zero value of int: %v, %T\n", zero, zero)

	// Interpreted string literals
	s1 := "Hello\nWorld\t very good!"
	fmt.Println(s1)

	// Raw string literals
	s2 := `H"ello
					\nWorld
	 very "good!`
	fmt.Println(s2)

	// Don't have type until be used
	const best_number = 42
	var number float32 = best_number // Now he's float
	fmt.Printf("best_number: %v, %T\n", best_number, best_number)
	fmt.Printf("number: %v, %T\n", number, number)

	// Successive values
	const (
		a = iota
		b
		_
		c
	)

	fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)

	const (
		d = iota + 3
		_
		e = iota * 10
	)

	fmt.Printf("d: %v, e: %v\n", d, e)
}
