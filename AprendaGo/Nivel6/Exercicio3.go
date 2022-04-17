package main

import "fmt"

func main() {

	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)

}

/*
- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
*/
