package main

import "fmt"

type pessoa struct {
	nome                    string
	sobrenome               string
	saboresFavoritosSorvete []string
}

func main() {
	pessoa1 := pessoa{
		"Lucas",
		"Freitas",
		[]string{"Chocolate", "Pistache"},
	}

	pessoa2 := pessoa{"Joao", "Silva", []string{"Morango"}}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa1.sobrenome)

	fmt.Println("\nSabores favoritos de sorvete:")
	for _, value := range pessoa1.saboresFavoritosSorvete {
		fmt.Println(value)
	}

	fmt.Println("")

	fmt.Println(pessoa2.nome)
	fmt.Println(pessoa2.sobrenome)

	fmt.Println("\nSabores favoritos de sorvete:")
	for _, value := range pessoa2.saboresFavoritosSorvete {
		fmt.Println(value)
	}

}

/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
