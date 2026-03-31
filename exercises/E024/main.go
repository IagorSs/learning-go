package main

import "fmt"

func main() {
	x := [][]string{
		{"Iagor", "Sousa", "Fight"},
		{"Mikki", "Godson", "Eat"},
		{"Lucci", "Morningstar", "Sleep"},
	}

	for _, v := range x {
		fmt.Println(v)
	}
}
