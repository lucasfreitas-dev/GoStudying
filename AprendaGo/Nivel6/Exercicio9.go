package main

import (
	"fmt"
)

func executaCallback(f func()) {
	f()
}

func main() {

	executaCallback(func() { fmt.Println("Printando via callback") })
}

/*
- Callback: passe uma função como argumento a outra função.
*/
