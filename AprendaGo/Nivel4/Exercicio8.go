package main

import "fmt"

func main() {
	hobbies := map[string][]string{
		"freitas_lucas": []string{"cerveja", "sair com amigos", "codar"},
		"silvar_joao":   []string{"futebol", "sair com amigos", "correr"},
		"silvar_jose":   []string{"volei", "filmes", "codar"},
	}

	for i, v := range hobbies {
		fmt.Printf("%v %v\n", i, v)
	}

}

/*
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
*/
