package main

import "fmt"

func main() {
	cadastro := struct {
		nome               string
		telefones          []string
		contatosEmergencia map[string]string
	}{
		nome:               "Lucas",
		telefones:          []string{"+55999-999-999"},
		contatosEmergencia: map[string]string{"maria": "+55888-888-888"},
	}

	fmt.Println(cadastro)

}

/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
