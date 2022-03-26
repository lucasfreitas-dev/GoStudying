package main

import "fmt"

func main() {
	nascimento := 1993
	atual := 2022

	for {
		fmt.Println(nascimento)

		if nascimento >= atual {
			break
		}
		nascimento++
	}
}

/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
