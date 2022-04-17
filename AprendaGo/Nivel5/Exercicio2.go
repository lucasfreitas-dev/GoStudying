package main

import "fmt"

type pessoa struct {
	nome                    string
	sobrenome               string
	saboresFavoritosSorvete []string
}

func main() {

	pessoas := make(map[string]pessoa)

	pessoas["Freitas"] = pessoa{
		"Lucas",
		"Freitas",
		[]string{"Chocolate", "Pistache"},
	}

	pessoas["Silva"] = pessoa{"Joao", "Silva", []string{"Morango"}}

	for _, pessoa := range pessoas {
		fmt.Println(pessoa.nome)
		fmt.Println(pessoa.sobrenome)

		fmt.Println("\nSabores favoritos")
		for _, sabor := range pessoa.saboresFavoritosSorvete {
			fmt.Println(sabor)
		}

		fmt.Println("")
	}

}

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/
