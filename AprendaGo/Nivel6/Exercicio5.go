package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) CalculaArea() float64 {
	return q.lado * q.lado
}

func (c circulo) CalculaArea() float64 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	CalculaArea() float64
}

func info(f figura) float64 {
	return f.CalculaArea()
}

func main() {
	circulo1 := circulo{2.1}
	quadrado1 := quadrado{6}

	fmt.Println(info(circulo1))
	fmt.Println(info(quadrado1))
}

/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/
