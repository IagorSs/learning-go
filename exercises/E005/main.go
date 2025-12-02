package main

import "fmt"

type myType int

var x myType

type mySubType myType

var y mySubType

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = mySubType(x)
	fmt.Printf("%v %T\n", y, y)
}
