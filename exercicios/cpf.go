package exercicios

import (
	"fmt"

	"github.com/Nhanderu/brdoc"
)

func Cpf() {

	var texto string

	fmt.Print("Informe um CPF para verificar se está na lista: ")
	fmt.Scan(&texto)

	if brdoc.IsCPF(texto) {
		fmt.Printf("%s é um CPF", texto)

	} else {
		fmt.Printf("%s nao é um CPF", texto)
	}
}
