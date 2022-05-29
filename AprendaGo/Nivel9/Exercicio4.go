package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count = 0

const numberRoutines = 200

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	wg.Add(numberRoutines)

	for i := 0; i < numberRoutines; i++ {
		go counter()
	}

	wg.Wait()
	fmt.Println(count)
}

func counter() {
	mu.Lock()
	countAux := count

	runtime.Gosched()

	countAux++
	count = countAux
	mu.Unlock()
	wg.Done()
}

/*
Utilize mutex para consertar a condição de corrida do exercício anterior.
*/
