package repository

import (
	"database/sql"
	"fmt"
	"minhascriptos/model"

	_ "github.com/lib/pq"
)

func AutenticarUsuario(email string) model.Usuario {

	db := StartDB()

	sqlStatement := `SELECT email FROM minhascriptosprincipal.usuario WHERE email=$1;`

	var usuario model.Usuario

	err := db.QueryRow(sqlStatement, email).Scan(&usuario.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada!")
		return model.Usuario{}
	case nil:
		return usuario
	default:
		return model.Usuario{}
	}
}

func CadastrarUsuario(usuario model.Usuario) int {

	db := StartDB()

	sqlStatement := `INSERT INTO minhascriptosprincipal.usuario(nome, email, senha) VALUES ($1, $2, $3) RETURNING id;`

	var id int

	err := db.QueryRow(sqlStatement, usuario.Nome, usuario.Email, usuario.Senha).Scan(&id)

	if err == nil {
		return id
	} else {
		return 0
	}
}

func ObterUsuario(email string, senha string) model.Usuario {

	db := StartDB()

	sqlStatement := `SELECT * FROM minhascriptosprincipal.usuario WHERE email=$1 AND senha=$2;`

	var usuario model.Usuario

	err := db.QueryRow(sqlStatement, email, senha).Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada!")
		return model.Usuario{}
	case nil:
		return usuario
	default:
		return model.Usuario{}
	}
}

func ObterDinheiroInserido(usuario_id int) []model.DinheiroInserido {

	db := StartDB()

	sqlStatement := `SELECT "tipoMoeda", sum("precoDeCompra")
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1
	GROUP BY "tipoMoeda"`

	var di model.DinheiroInserido
	var dis []model.DinheiroInserido

	rows, err := db.Query(sqlStatement, usuario_id)

	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Println("Row não encontrada!")
			return []model.DinheiroInserido{}

		} else {
			fmt.Println("Erro ao obter dinheiro inserido | Erro:", err)
			return []model.DinheiroInserido{}
		}

	}

	for rows.Next() {

		err = rows.Scan(&di.TipoMoeda, &di.Total)

		if err != nil {
			fmt.Println("Erro ao obter dinheiro inserido | Erro ao fazer Scan da row", err)
			return []model.DinheiroInserido{}
		}

		dis = append(dis, di)
	}

	err = rows.Err()

	if err != nil {
		return []model.DinheiroInserido{}
	}

	return dis
}

func ObterUsuarioPeloID(usuario_id int) int {

	db := StartDB()

	sqlStatement := `SELECT id FROM minhascriptosprincipal.usuario WHERE id=$1`

	var id int

	err := db.QueryRow(sqlStatement, usuario_id).Scan(&id)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada!")
		return 0
	case nil:
		return id
	default:
		return 0
	}
}
