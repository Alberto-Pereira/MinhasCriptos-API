package repository

import (
	"database/sql"
	"fmt"
	"minhascriptos/model"
)

func AdicionarMoeda(cripto model.Cripto) int {
	db := StartDB()

	sqlStatement := `INSERT INTO minhascriptosprincipal.criptomoeda(
					"tipoMoeda", "dataDeCompra", "quantidadeComprada", "precoDeCompra", "valorDaUnidadeNoDiaDeCompra", 
					usuario_id)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id;`

	var id int

	reqRow := db.QueryRow(sqlStatement, cripto.TipoMoeda, cripto.DataDeCompra, cripto.QuantidadeComprada,
		cripto.PrecoDeCompra, cripto.ValorDaUnidadeNoDiaDeCompra, cripto.UsuarioId.ID).Scan(&id)

	switch reqRow {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada")
	case nil:
		fmt.Println("Moeda adicionada com ID:", id)
		return id
	default:
		panic(reqRow)
	}

	return id
}

func EditarMoeda(cripto model.Cripto) int {
	db := StartDB()

	sqlStatement := `UPDATE minhascriptosprincipal.criptomoeda
	SET "tipoMoeda"=$2, "dataDeCompra"=$3, "quantidadeComprada"=$4, "precoDeCompra"=$5, 
	"valorDaUnidadeNoDiaDeCompra"=$6
	WHERE id=$1 AND usuario_id=$7
	RETURNING id, "tipoMoeda";`

	var id int
	var tipoMoeda string

	reqRow := db.QueryRow(sqlStatement, cripto.ID, cripto.TipoMoeda, cripto.DataDeCompra,
		cripto.QuantidadeComprada, cripto.PrecoDeCompra, cripto.ValorDaUnidadeNoDiaDeCompra, cripto.UsuarioId.ID).
		Scan(&id, &tipoMoeda)

	switch reqRow {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada")
	case nil:
		fmt.Printf("Cripto com ID: %v e Tipo: %v foi EDITADA!\n", id, tipoMoeda)
	default:
		panic(reqRow)
	}

	return id
}

func DeletarMoeda(cripto model.Cripto) int {
	db := StartDB()

	sqlStatement := `DELETE FROM minhascriptosprincipal.criptomoeda 
	WHERE id=$1 and usuario_id=$2
	RETURNING id;`

	var id int

	reqRow := db.QueryRow(sqlStatement, cripto.ID, cripto.UsuarioId.ID).Scan(&id)

	switch reqRow {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada!")
	case nil:
		fmt.Println("ID para DELETAR:", cripto.ID, " | ID deletado:", id)
		return id
	default:
		panic(reqRow)
	}

	return id
}

func ObterMoedas(usuario_id int) []model.Cripto {
	db := StartDB()

	sqlStatement := `SELECT id, "tipoMoeda", "dataDeCompra", "quantidadeComprada", "precoDeCompra", 
	"valorDaUnidadeNoDiaDeCompra", usuario_id
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1;`

	rows, err := db.Query(sqlStatement, usuario_id)

	if err != nil {
		fmt.Println("Obter Moedas | Erro ao realizar Query:", err)
		panic(err)
	}
	defer db.Close()

	var cripto model.Cripto
	var criptos []model.Cripto

	for rows.Next() {
		err = rows.Scan(&cripto.ID, &cripto.TipoMoeda, &cripto.DataDeCompra, &cripto.QuantidadeComprada,
			&cripto.PrecoDeCompra, &cripto.ValorDaUnidadeNoDiaDeCompra, &cripto.UsuarioId.ID)

		if err != nil {
			fmt.Println("Obter moedas | Erro ao realizar Scan:", err)
			panic(err)
		}

		fmt.Println(cripto)
		criptos = append(criptos, cripto)
	}

	fmt.Println(criptos)

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return criptos
}

func ObterMoedasBuscaPersonalizada(usuario_id int, tipoMoeda string, dataDeCompra string) []model.Cripto {
	db := StartDB()

	sqlStatement := `SELECT id, "tipoMoeda", "dataDeCompra", "quantidadeComprada", "precoDeCompra", 
	"valorDaUnidadeNoDiaDeCompra", usuario_id
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1 AND ("tipoMoeda"=$2 OR "dataDeCompra"=$3);`

	rows, err := db.Query(sqlStatement, usuario_id, tipoMoeda, dataDeCompra)

	if err != nil {
		fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Query:", err)
		panic(err)
	}
	defer db.Close()

	var cripto model.Cripto
	var criptos []model.Cripto

	for rows.Next() {
		err = rows.Scan(&cripto.ID, &cripto.TipoMoeda, &cripto.DataDeCompra, &cripto.QuantidadeComprada,
			&cripto.PrecoDeCompra, &cripto.ValorDaUnidadeNoDiaDeCompra, &cripto.UsuarioId.ID)

		if err != nil {
			fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Scan:", err)
			panic(err)
		}

		fmt.Println(cripto)
		criptos = append(criptos, cripto)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return criptos
}
