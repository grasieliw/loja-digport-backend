package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() {

	CriaTabelas()
}

func ConectaBancoDados() *sql.DB {
	dbPass := os.Getenv("DB_PASS")
	//connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	connStr := fmt.Sprint("user=postgres dbname=digport_loja password=", dbPass, " host=localhost sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CriaTabelas() {
	db := ConectaBancoDados()
	defer db.Close()

	createUsersTable := `
CREATE TABLE IF NOT EXISTS USUARIO (
	id SERIAL PRIMARY KEY,
	nome VARCHAR,
	telefone VARCHAR,
	endereco VARCHAR,
	email VARCHAR NOT NULL UNIQUE,
	senha VARCHAR NOT NULL
);`
	_, err := db.Exec(createUsersTable)

	if err != nil {
		panic("Erro ao criar tabela usuario")
	} else {
		fmt.Println("Tabela usuario criada")
	}

	createProductsTableScript := `
    CREATE TABLE IF NOT EXISTS produtos
    (
    id varchar primary key,
    nome varchar,
    preco  float8,
    descricao varchar,
    imagem varchar,
    quantidade int
    );`

	_, err = db.Exec(createProductsTableScript)

	if err != nil {
		panic("Erro ao criar tabela usuario")
	} else {
		fmt.Println("Tabela usuario criada")
	}

}
