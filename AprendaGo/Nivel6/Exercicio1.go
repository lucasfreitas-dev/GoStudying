package main

import "fmt"

func returnInt() int {
	return 1
}

func returnIntEString() (int, string) {
	return 1, "um"
}

func main() {

	fmt.Println(returnInt())
	fmt.Println(returnIntEString())

}

/*
- Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/
