package main

import "fmt"

func main() {
	nascimento := 1993
	atual := 2022

	for nascimento <= atual {
		fmt.Println(nascimento)
		nascimento++
	}
}

/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/
