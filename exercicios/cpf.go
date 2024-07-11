package exercicios

import (
	"fmt"

	"github.com/Nhanderu/brdoc"
)

func Cpf() {

	var texto string
	var isValidCpf bool
	isValidCpf = brdoc.IsCPF(texto)

	//caso sim
	fmt.Printf("%s é um CPF", texto)
	//caso nao
	fmt.Printf("%s nao é um CPF", texto)
}
