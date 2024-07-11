package exercicios

import "fmt"

func Exercicio1() {
	minhasDespesas := []string{"patins", "manutencao", "aulas", "academia"}

	fmt.Println("Minha lista de despesas de patinacao tem %d itens", len(minhasDespesas))

	for _, despesa := range minhasDespesas {
		fmt.Println(despesa)
	}

	fmt.Println("Qual item voce deseja saber se esta presenta na nossa lista de despesas de patinacao?")

	var itemVerificar string
	fmt.Scan(&itemVerificar)

	fmt.Println("Verificando se há", itemVerificar, "na sua lista...")
	for _, v := range minhasDespesas {

		if v == itemVerificar {
			fmt.Println("Sim")
			return
		}
	}
	fmt.Println("Nao")
}
