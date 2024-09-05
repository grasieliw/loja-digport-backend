package main

import (
	"github.com/grasieliw/loja-digport-backend/db"
	_ "github.com/lib/pq"
)

func main() {
	db.InitDB()
	StarServer()

	//fmt.Println("Bem Vindo(a) à loja DigPort")

}
