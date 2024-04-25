package main

import "fmt"

func main() {
	StarServer()
	produtosFiltrados := buscaPorNome("Agua")
	fmt.Println(produtosFiltrados)

}
