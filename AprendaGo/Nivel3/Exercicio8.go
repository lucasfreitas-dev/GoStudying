package main

import "fmt"

func main() {
	x := 5

	switch {
	case x == 42:
		fmt.Println("x é 42!")
	case x == 10:
		fmt.Println("x é 10!")
	case x == 5:
		fmt.Println("x é 5!")
	}
}

// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
