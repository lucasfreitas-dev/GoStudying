package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "basquete":
		fmt.Println("meu esporte favorito é basquete")
	case "volei":
		fmt.Println("meu esporte favorito é volei")
	case "futebol":
		fmt.Println("meu esporte favorito é futebol")
	}
}

// - Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
