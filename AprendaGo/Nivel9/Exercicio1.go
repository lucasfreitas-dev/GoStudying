package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go routineOne()
	go routineTwo()

	wg.Wait()
}

func routineOne() {
	fmt.Println("One!")
	wg.Done()
}

func routineTwo() {
	fmt.Println("Two!")
	wg.Done()
}

/*
Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/
