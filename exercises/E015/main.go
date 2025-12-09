package main

import (
	"fmt"
	"time"
)

func main() {
	year := 2000
	currentYear := time.Now().Year()

	for year <= currentYear {
		fmt.Println(year)
		year++
	}
}
