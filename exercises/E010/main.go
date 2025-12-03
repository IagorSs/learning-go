package main

import "fmt"

func main() {
	x := `Hello,
This is a raw string literal.
It can span multiple lines and include "quotes" without needing to escape them.
Goodbye!`
	fmt.Println(x)
}
