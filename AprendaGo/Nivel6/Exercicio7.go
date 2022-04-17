package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("Printando de dentro da função")
	}

	x()
}

/*
- Atribua uma função a uma variável.
- Chame essa função.
*/
