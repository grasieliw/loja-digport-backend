package exercicios

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

func Exercicio3() {
	contacts := make(map[string]Contact)
	//add contatos no mapa
	contacts["Pedro"] = Contact{Name: "Pedro Silva", Phone: "51998784236", Email: "pedro.silva@gmail.com"}
	contacts["Maria"] = Contact{Name: "Maria Silva", Phone: "51998654123", Email: "maria.silva@gmail.com"}

	//imprimindo contatos
	fmt.Println("Contatos já existentes:")
	for _, contact := range contacts {
		fmt.Println("Nome:", contact.Name)
		fmt.Println("Phone:", contact.Phone)
		fmt.Println("Email:", contact.Email)
	}

	var get_contato string
	fmt.Printf("Digite o nome do contato que deseja encontrar: ")
	fmt.Scan(&get_contato)

	//procurando contatos
	v, encontrada := contacts[get_contato]
	if encontrada {
		fmt.Printf("Os detalhes do contato %s informado é: %s\n", get_contato, v)
	} else {
		fmt.Printf("Contato %s não encontrado\n", get_contato)
	}
}
