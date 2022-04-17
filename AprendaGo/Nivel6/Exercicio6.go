package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Printando de dentro da função")
	}()
}

/*
- Crie e utilize uma função anônima.
*/
