// Package repository contém as operações de repositório das entidades usuário e cripto
// Contém também a configuração do banco de dados
package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 123
	dbname   = "minhascriptos"
)

// Start DB
// Iniciliza o banco de dados com as constantes informadas
// Retorna o banco de dados inicializado
func StartDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%d dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Banco de dados conectado!")

	return db
}
