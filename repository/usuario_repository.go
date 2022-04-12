package repository

import (
	"database/sql"
	"fmt"
	"minhascriptos/model"

	_ "github.com/lib/pq"
)

func AutenticarUsuario(email string) model.Usuario {

	db := StartDB()

	sqlStatement := `SELECT * FROM minhascriptosprincipal.usuario WHERE email=$1;`

	var usuario model.Usuario

	reqRow := db.QueryRow(sqlStatement, email).Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha)

	switch reqRow {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada!")
	case nil:
		fmt.Println(usuario)
	default:
		panic(reqRow)
	}

	return usuario
}

func CadastrarUsuario(usuario model.Usuario) (id int) {
	db := StartDB()

	sqlStatement := `INSERT INTO minhascriptosprincipal.usuario(nome, email, senha) VALUES ($1, $2, $3) RETURNING id;`

	id = 0

	reqRow := db.QueryRow(sqlStatement, usuario.Nome, usuario.Email, usuario.Senha).Scan(&id)

	switch reqRow {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada!")
	case nil:
		fmt.Println(id)
	default:
		panic(reqRow)
	}

	return id
}

func ObterUsuario(email string, senha string) model.Usuario {

	db := StartDB()

	sqlStatement := `SELECT * FROM minhascriptosprincipal.usuario 
					WHERE email=$1 AND senha=$2;`

	var usuario model.Usuario

	reqRow := db.QueryRow(sqlStatement, email, senha).Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha)

	switch reqRow {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada!")
	case nil:
		fmt.Println(usuario)
		return usuario
	default:
		panic(reqRow)
	}

	return usuario
}

func ObterDinheiroInserido(usuario_id int) model.DinheiroInseridoSlice {
	db := StartDB()

	sqlStatement := `SELECT "tipoMoeda", sum("precoDeCompra")
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1
	GROUP BY "tipoMoeda"`

	var di model.DinheiroInserido
	var dis model.DinheiroInseridoSlice

	rows, err := db.Query(sqlStatement, usuario_id)

	if err != nil {
		fmt.Println("Erro ao obter dinheiro inserido | Erro:", err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&di.TipoMoeda, &di.Total)

		if err != nil {
			fmt.Println("Erro ao obter dinheiro inserido | Erro ao fazer Scan da row", err)
			panic(err)
		}

		fmt.Println(di.TipoMoeda, di.Total)
		dis = append(dis, di)
	}

	fmt.Println(dis)

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return dis
}
