package main

import "fmt"

func main() {
	x := 10

	if x == 42 {
		fmt.Println("x é 42!")
	} else if x == 10 {
		fmt.Println("na verdade x é 10!")
	} else {
		fmt.Println("x não é 10 nem 42")
	}
}

// - Utilizando a solução anterior, adicione as opções else if e else.
