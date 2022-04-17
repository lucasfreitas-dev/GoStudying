package main

import (
	"fmt"
)

func retornaOutraFuncao() func() string {
	x := "AQUI"
	return func() string {
		return "Estou dentro da função, e consigo acessar uma variavel do escopo acima da meu via closure:" + x
	}
}

func main() {
	x := retornaOutraFuncao()

	fmt.Println(x())
}

/*
- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/
