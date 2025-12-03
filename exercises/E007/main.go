package main

import "fmt"

func main() {
	lowest := 56.77
	highest := 99.99

	fmt.Println("Comparison between numbers", lowest, highest)

	equality := lowest == highest
	fmt.Println("Equal?", equality)

	inequality := lowest != highest
	fmt.Println("Not equal?", inequality)

	less_equal := lowest <= highest
	fmt.Println("Less or equal?", less_equal)

	less := lowest < highest
	fmt.Println("Less?", less)

	greater_equal := lowest >= highest
	fmt.Println("Greater or equal?", greater_equal)

	greater := lowest > highest
	fmt.Println("Greater?", greater)
}
