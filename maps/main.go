package main

import "fmt"

func main() {
	friends := map[string]int{
		"alfredo": 123,
		"joana":   999,
	}

	fmt.Println(friends)
	fmt.Println(friends["alfredo"])

	// Add values on map
	friends["gopher"] = 111
	fmt.Println(friends["gopher"])

	// Go returns 0 if some value isn't found on map, to distinguish this case
	// to real 0 value a second variable is returned with this info
	ghost, ghost_found := friends["ghost"]
	fmt.Println(ghost, ghost_found)

	// A reduced way to do some thing related to not found case (comma ok idiom)
	if jorge, jorge_found := friends["jorge"]; !jorge_found {
		fmt.Println("Cannot found jorge!")
	} else {
		fmt.Println(jorge)
	}

	// Order isn't guaranteed and still can travel him with range
	for k, v := range friends {
		fmt.Println(k, v)
	}
}
