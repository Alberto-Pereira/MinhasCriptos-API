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
		return 0
	case nil:
		return id
	default:
		return 0
	}
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
		return 0
	case nil:
		return id
	default:
		return 0
	}
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
		return 0
	case nil:
		return id
	default:
		return 0
	}
}

func ObterMoedas(usuario_id int) []model.Cripto {

	db := StartDB()

	sqlStatement := `SELECT id, "tipoMoeda", to_char("dataDeCompra", 'YYYY-MM-DD'), "quantidadeComprada", "precoDeCompra", 
	"valorDaUnidadeNoDiaDeCompra", usuario_id
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1;`

	rows, err := db.Query(sqlStatement, usuario_id)

	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Println("Row não encontrada!")
			return []model.Cripto{}

		} else {
			fmt.Println("Obter Moedas | Erro ao realizar Query:", err)
			return []model.Cripto{}
		}
	}

	// defer db.Close() TESTAR

	var cripto model.Cripto
	var criptos []model.Cripto

	for rows.Next() {
		err = rows.Scan(&cripto.ID, &cripto.TipoMoeda, &cripto.DataDeCompra, &cripto.QuantidadeComprada,
			&cripto.PrecoDeCompra, &cripto.ValorDaUnidadeNoDiaDeCompra, &cripto.UsuarioId.ID)

		if err != nil {
			fmt.Println("Obter moedas | Erro ao realizar Scan:", err)
			return []model.Cripto{}
		}

		criptos = append(criptos, cripto)
	}

	err = rows.Err()

	if err != nil {
		return []model.Cripto{}
	}

	return criptos
}

func ObterMoedasBuscaPersonalizada(usuario_id int, tipoMoeda string, dataDeCompra string) []model.Cripto {

	if tipoMoeda != "" && dataDeCompra != "" {
		moedas := obterMoedaPorTipoMoedaEDataDeCompra(usuario_id, tipoMoeda, dataDeCompra)
		return moedas
	} else if tipoMoeda != "" {
		moedas := obterMoedaPorTipoMoeda(usuario_id, tipoMoeda)
		return moedas
	} else {
		moedas := obterMoedaPorDataDeCompra(usuario_id, dataDeCompra)
		return moedas
	}

}

func obterMoedaPorTipoMoedaEDataDeCompra(usuario_id int, tipoMoeda string, dataDeCompra string) []model.Cripto {

	db := StartDB()

	sqlStatement := `SELECT id, "tipoMoeda", to_char("dataDeCompra", 'YYYY-MM-DD'), "quantidadeComprada", "precoDeCompra", 
	"valorDaUnidadeNoDiaDeCompra", usuario_id
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1 AND "tipoMoeda"=$2 AND "dataDeCompra"=$3;`

	rows, err := db.Query(sqlStatement, usuario_id, tipoMoeda, dataDeCompra)

	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Println("Row não encontrada!")
			return []model.Cripto{}
		} else {
			fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Query:", err)
			return []model.Cripto{}
		}
	}

	//defer db.Close()

	var cripto model.Cripto
	var criptos []model.Cripto

	for rows.Next() {
		err = rows.Scan(&cripto.ID, &cripto.TipoMoeda, &cripto.DataDeCompra, &cripto.QuantidadeComprada,
			&cripto.PrecoDeCompra, &cripto.ValorDaUnidadeNoDiaDeCompra, &cripto.UsuarioId.ID)

		if err != nil {
			fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Scan:", err)
			return []model.Cripto{}
		}

		criptos = append(criptos, cripto)
	}

	err = rows.Err()

	if err != nil {
		return []model.Cripto{}
	}

	return criptos
}

func obterMoedaPorTipoMoeda(usuario_id int, tipoMoeda string) []model.Cripto {

	db := StartDB()

	sqlStatement := `SELECT id, "tipoMoeda", to_char("dataDeCompra", 'YYYY-MM-DD'), "quantidadeComprada", "precoDeCompra", 
	"valorDaUnidadeNoDiaDeCompra", usuario_id
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1 AND "tipoMoeda"=$2;`

	rows, err := db.Query(sqlStatement, usuario_id, tipoMoeda)

	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Println("Row não encontrada!")
			return []model.Cripto{}
		} else {
			fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Query:", err)
			return []model.Cripto{}
		}
	}

	//defer db.Close()

	var cripto model.Cripto
	var criptos []model.Cripto

	for rows.Next() {
		err = rows.Scan(&cripto.ID, &cripto.TipoMoeda, &cripto.DataDeCompra, &cripto.QuantidadeComprada,
			&cripto.PrecoDeCompra, &cripto.ValorDaUnidadeNoDiaDeCompra, &cripto.UsuarioId.ID)

		if err != nil {
			fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Scan:", err)
			return []model.Cripto{}
		}

		criptos = append(criptos, cripto)
	}

	err = rows.Err()

	if err != nil {
		return []model.Cripto{}
	}

	return criptos
}

func obterMoedaPorDataDeCompra(usuario_id int, dataDeCompra string) []model.Cripto {

	db := StartDB()

	sqlStatement := `SELECT id, "tipoMoeda", to_char("dataDeCompra", 'YYYY-MM-DD'), "quantidadeComprada", "precoDeCompra", 
	"valorDaUnidadeNoDiaDeCompra", usuario_id
	FROM minhascriptosprincipal.criptomoeda
	WHERE usuario_id=$1 AND "dataDeCompra"=$2;`

	rows, err := db.Query(sqlStatement, usuario_id, dataDeCompra)

	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Println("Row não encontrada!")
			return []model.Cripto{}
		} else {
			fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Query:", err)
			return []model.Cripto{}
		}
	}

	//defer db.Close()

	var cripto model.Cripto
	var criptos []model.Cripto

	for rows.Next() {
		err = rows.Scan(&cripto.ID, &cripto.TipoMoeda, &cripto.DataDeCompra, &cripto.QuantidadeComprada,
			&cripto.PrecoDeCompra, &cripto.ValorDaUnidadeNoDiaDeCompra, &cripto.UsuarioId.ID)

		if err != nil {
			fmt.Println("Obter Moedas Busca Personalizada | Erro ao realizar Scan:", err)
			return []model.Cripto{}
		}

		criptos = append(criptos, cripto)
	}

	err = rows.Err()

	if err != nil {
		return []model.Cripto{}
	}

	return criptos
}

func ObterMoedaPeloID(cripto model.Cripto) int {

	db := StartDB()

	sqlStatement := `SELECT id FROM minhascriptosprincipal.criptomoeda WHERE id=$1 AND usuario_id=$2`

	var id int

	err := db.QueryRow(sqlStatement, cripto.ID, cripto.UsuarioId.ID).Scan(&id)

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
