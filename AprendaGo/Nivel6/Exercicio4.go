package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) GetNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func main() {

	lucas := pessoa{"Lucas", "Freitas", 25}

	fmt.Println(lucas.GetNomeCompleto())
}

/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/
