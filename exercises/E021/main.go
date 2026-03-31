package main

import "fmt"

var x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

func printSlice() {
	fmt.Println(x)
}

func main() {
	printSlice()

	x = append(x, 52)
	printSlice()

	x = append(x, 53, 54, 55)
	printSlice()

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	printSlice()
}
