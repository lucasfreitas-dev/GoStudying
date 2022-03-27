package main

import "fmt"

func main() {
	cadastro := [][]string{
		[]string{
			"Lucas", "Freitas", "cerveja",
		},
		[]string{
			"João", "Silva", "futebol",
		},
		[]string{
			"José", "Silva", "volei",
		},
	}

	for i := range cadastro {
		for j := range cadastro {
			fmt.Printf("%v ", cadastro[i][j])
		}
		fmt.Print("\n")
	}

}

/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
