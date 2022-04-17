package main

import "fmt"

type veiculo struct {
	cor    string
	portas int
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

func main() {

	hilux := caminhonete{
		veiculo: veiculo{
			cor:    "preto",
			portas: 4,
		},
		tracaoNasQuatro: true,
	}

	civic := sedan{veiculo{"branco", 4}, false}

	fmt.Println(hilux)
	fmt.Println(hilux.cor)

	fmt.Println("")

	fmt.Println(civic)
	fmt.Println(civic.cor)

}

/*
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
*/
