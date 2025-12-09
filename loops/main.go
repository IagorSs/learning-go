package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	x := 1
	for x < 10 {
		if x%2 == 0 {
			x++
			continue
		}

		fmt.Println("Odd: ", x)
		x++
	}

	x = 1
	for {
		if x > 5 {
			fmt.Println("Breaking the eternal loop")
			break
		}

		fmt.Println("Eternal loop")
		x++
	}
}
