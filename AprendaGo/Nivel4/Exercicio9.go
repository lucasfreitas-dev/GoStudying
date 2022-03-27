package main

import "fmt"

func main() {
	hobbies := map[string][]string{
		"freitas_lucas": []string{"cerveja", "sair com amigos", "codar"},
		"silva_joao":    []string{"futebol", "sair com amigos", "correr"},
		"silva_jose":    []string{"volei", "filmes", "codar"},
	}

	hobbies["souza_leandro"] = []string{"esportes", "dança", "estudar"}

	for i, v := range hobbies {
		fmt.Printf("%v %v\n", i, v)
	}

}

// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
