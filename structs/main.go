package main

import "fmt"

type person struct {
	name    string
	surname string
}

type restaurant_client struct {
	person
	smoker bool
}

func main() {
	person1 := person{
		name:    "Albert",
		surname: "Einstein",
	}

	// Omitting params names
	person2 := person{"Nikolas", "Tesla"}

	fmt.Println(person1, person2)

	client1 := restaurant_client{
		person: person1,
		smoker: true,
	}

	fmt.Println(client1)

	// Can access fields of internal structs in external if don't have conflicts
	fmt.Println(client1.name, client1.smoker)
}
