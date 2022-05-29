package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var count int64 = 0

const numberRoutines = 200

var wg sync.WaitGroup

func main() {
	wg.Add(numberRoutines)

	for i := 0; i < numberRoutines; i++ {
		go counter()
	}

	wg.Wait()
	fmt.Println(count)
}

func counter() {

	atomic.AddInt64(&count, 1)

	runtime.Gosched()

	wg.Done()
}

/*
Utilize atomic para consertar a condição de corrida do exercício #3.
*/
