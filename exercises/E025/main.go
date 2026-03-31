package main

import "fmt"

func main() {
	persons := map[string][]string{
		"Sousa_Iagor":       {"Fight", "Code", "Read"},
		"Godson_Mikki":      {"Eat", "Toy", "Destroy things"},
		"Morningstar_Lucci": {"Sleep", "Follow lights", "Hunt"},
	}

	for name, hobbies := range persons {
		fmt.Println("----------------------------")
		fmt.Println(name)

		for i, hobbie := range hobbies {
			fmt.Printf("\tHobbie %d: %s\n", i, hobbie)
		}
	}

	fmt.Println("----------------------------")
}
