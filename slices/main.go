package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Can't do it with array because is fixed
	// array2 := append(array, 6)
	slice2 := append(slice, 6)
	fmt.Println(slice2)

	// Throw panic setting index out of range
	// slice[20] = 1

	fruits := []string{"banana", "apple", "lemon"}
	// How travel the entire slice
	for i, v := range fruits {
		fmt.Println("No índice", i, "temos a fruta:", v)
	}
	// How ignore some prop
	for _, v := range fruits {
		fmt.Println("Uma das frutas é", v)
	}

	flavors := []string{"Margherita", "4cheese", "portuguese", "house"}
	// Slicing the slice, will includes the start index and exclude the final index
	pizza_slice := flavors[1:3]
	fmt.Println(pizza_slice)

	// Reduced form to slice a slice, follow the rule to include start and exclude final
	pizza_slice2 := flavors[1:]
	fmt.Println(pizza_slice2)
	pizza_slice3 := flavors[:2]
	fmt.Println(pizza_slice3)

	// Deep copy of slice without loops
	pizza_slice_copy := pizza_slice[:]
	fmt.Println(pizza_slice_copy)

	// How append all elements from some slice to another
	pizza_slice23 := append(pizza_slice2, pizza_slice3...)
	fmt.Println(pizza_slice23)

	// How delete element in slice
	flavors_without_cheese := append(flavors[:1], flavors[2:]...)
	fmt.Println(flavors_without_cheese)

	// Slices are made by array, you can check the "capacity" of current array related to slice
	elements := []int{1, 2, 3}
	fmt.Println(len(elements), cap(elements))

	// When the capacity is exceeded the array need to be remade with new len
	elements = append(elements, 4)
	fmt.Println(len(elements), cap(elements))

	// New cap is double from previous
	elements = append(elements, 5, 6, 7)
	fmt.Println(len(elements), cap(elements))

	// Make turns more controllable initial array and capacity of slice
	made_elements := make([]int, 5, 10)
	made_elements[0], made_elements[1], made_elements[2] = 1, 2, 3
	fmt.Println(made_elements, len(made_elements), cap(made_elements))

	multidimensional_slice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(multidimensional_slice)

	// The problem with subjacent array in slices:
	fmt.Println(flavors)
	// Explanation: when I'm doing flavors_without_cheese I use append using ref
	// to flavors slice, so the operation have be done in subjacent array and edit
	// other slices that are manipulating same guy
	// Recommendation to avoid this: use just one variable for each array reference
}
