package main

import "fmt"

func main() {
	s := "Olá, Go!"
	fmt.Println(s)

	sb := []byte(s)
	fmt.Printf("%v\n%T\n", sb, sb)

	for _, v := range sb {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}
}
