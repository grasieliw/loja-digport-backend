package exercicios

import "fmt"

func Exercicio2() {
	despesas := map[string]int{"aluguel": 1800, "internet": 100, "eletricidade": 50}

	total := 0

	var get_despesa string
	fmt.Printf("digite a despesa: ")
	fmt.Scan(&get_despesa)

	v, encontrada := despesas[get_despesa]
	if encontrada {
		fmt.Printf("O valor da despesa %s informada é: %d\n", get_despesa, v)
	} else {
		fmt.Printf("Despesa %s não encontrada\n", get_despesa)
	}

	for despesa, valor := range despesas {

		total = total + valor
		fmt.Printf("A despesa de %s é %d\n", despesa, valor)
	}

	orcamento := 0
	fmt.Printf("Digite o seu orçamento: ")
	fmt.Scan(&orcamento)

	if total > orcamento {
		fmt.Printf("O total das despesas %d ultrapassou o seu orçamento %d \n", total, orcamento)
	}
	fmt.Printf("O total das despesas é %d\n", total)
}
