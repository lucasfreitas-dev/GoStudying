package main

import (
	"fmt"
)

func retornaOutraFuncao() func() string {
	return func() string {
		return "Estou dentro da função"
	}
}

func main() {
	x := retornaOutraFuncao()

	fmt.Println(x())
}

/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
*/
