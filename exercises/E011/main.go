package main

import "fmt"

func main() {
	const (
		year = iota + 2026
		year1
		year2
		year3
	)

	fmt.Println(year, year1, year2, year3)
}
