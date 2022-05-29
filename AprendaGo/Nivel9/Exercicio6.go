package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
}

/*
Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build
    - go install
*/
