package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count = 0

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
	countAux := count

	runtime.Gosched()

	countAux++
	count = countAux

	wg.Done()
}

/*
Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
*/
