package main

import "fmt"

func main() {
	brazillian_states := make([]string, 26)
	fmt.Println(brazillian_states, len(brazillian_states), cap(brazillian_states))

	copy(brazillian_states, []string{
		"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins",
	})
	fmt.Println(brazillian_states, len(brazillian_states), cap(brazillian_states))
}
