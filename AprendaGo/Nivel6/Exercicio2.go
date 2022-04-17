package main

import "fmt"

func soma(itens ...int) int {

	soma := 0

	for _, value := range itens {
		soma += value
	}

	return soma
}

func somaSlice(itens []int) int {

	soma := 0

	for _, value := range itens {
		soma += value
	}

	return soma
}

func main() {

	itens := []int{1, 2, 3}

	fmt.Println(soma(itens...))
	fmt.Println(somaSlice(itens))

}

/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/
