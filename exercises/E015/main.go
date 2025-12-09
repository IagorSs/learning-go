package main

import (
	"fmt"
	"time"
)

func main() {
	birthYear := 2000
	currentYear := time.Now().Year()

	for year := birthYear; year <= currentYear; year++ {
		fmt.Println(year)
	}
}
