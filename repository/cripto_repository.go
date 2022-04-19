// Package repository contém as operações de repositório das entidades usuário e cripto
// Contém também a configuração do banco de dados
package repository

import (
	"database/sql"
	"fmt"
	"minhascriptos/model"
)

// Adicionar Moeda
// A entidade cripto é a moeda a ser enviada para o banco de dados para ser adicionada
// Se ela for adicionada, é retornado o id correspondente
// Se ela não for adicionada, é retornado 0
func AdicionarMoeda(cripto model.Cripto) int {

	db := StartDB()

	sqlStatement := `INSERT INTO minhascriptosprincipal.criptomoeda(
					"tipoMoeda", "dataDeCompra", "quantidadeComprada", "precoDeCompra", "valorDaUnidadeNoDiaDeCompra", 
					usuario_id)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id;`

	var id int

	err := db.QueryRow(sqlStatement, cripto.TipoMoeda, cripto.DataDeCompra, cripto.QuantidadeComprada,
		cripto.PrecoDeCompra, cripto.ValorDaUnidadeNoDiaDeCompra, cripto.UsuarioId.ID).Scan(&id)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada")
		return 0
	case nil:
		return id
	default:
		return 0
	}
}

// Editar Moeda
// A entidade cripto é a moeda a ser enviada para o banco de dados para ser atualizada
// Se ela for atualizada, é retornado o id correspondente
// Se ela não for atualizada, é retornado 0
func EditarMoeda(cripto model.Cripto) int {

	db := StartDB()

	sqlStatement := `UPDATE minhascriptosprincipal.criptomoeda
	SET "tipoMoeda"=$2, "dataDeCompra"=$3, "quantidadeComprada"=$4, "precoDeCompra"=$5, 
	"valorDaUnidadeNoDiaDeCompra"=$6
	WHERE id=$1 AND usuario_id=$7
	RETURNING id;`

	var id int

	err := db.QueryRow(sqlStatement, cripto.ID, cripto.TipoMoeda, cripto.DataDeCompra,
		cripto.QuantidadeComprada, cripto.PrecoDeCompra, cripto.ValorDaUnidadeNoDiaDeCompra, cripto.UsuarioId.ID).
		Scan(&id)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Row não encontrada")
		return 0
	case nil:
		return id
	default:
		return 0
	}
}

// Deletar Moeda
// A entidade cripto é a moeda a ser enviada para o banco de dados para ser deletada
// Se ela for deletada, é retornado o id correspondente
// Se ela não for deletada, é retornado 0
func DeletarMoeda(cripto model.Cripto) int {

	db := StartDB()

	sqlStatement := `DELETE FROM minhascriptosprincipal.criptomoeda 
	WHERE id=$1 and usuario_id=$2
	RETURNING id;`

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

// Obter Moedas
// É fornecido o id de um usuário, a ser enviado para o banco de dados, para obter as moedas correspondentes
// Se o usuário existir e tiver moedas registradas, é retornado as moedas correspondentes
// Se o usuário existir e não tiver moedas registradas, é retornado nil
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

// Obter Moedas Busca Personalizada
// É fornecido o id de um usuário com parâmetros personalizados, a serem enviados para o banco de dados,
// para obter as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas corresponderem aos parâmetros, é retornado as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas não corresponderem aos parâmetros, é retornado nil
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

// Obter Moeda Por Tipo Moeda E Data De Compra
// É fornecido o id de um usuário com os parâmetro de tipo de moeda e data de compra, que serão enviados para o banco de dados,
// para obter as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas corresponderem aos parâmetros, é retornado as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas não corresponderem aos parâmetros, é retornado nil
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

// Obter Moeda Por Tipo Moeda
// É fornecido o id de um usuário e o parâmetro de tipo de moeda, que serão enviados para o banco de dados,
// para obter as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas corresponderem ao parâmetro, é retornado as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas não corresponderem ao parâmetro, é retornado nil
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

// Obter Moeda Por  Data De Compra
// É fornecido o id de um usuário e o parâmetro de data de compra, que serão enviados para o banco de dados,
// para obter as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas corresponderem ao parâmetro, é retornado as moedas correspondentes
// Se o usuário existir, tiver moedas registradas e elas não corresponderem ao parâmetro, é retornado nil
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

// Obter Moeda Pelo ID
// É fornecido o id de um usuário e o id de uma moeda, utilizando uma entidade cripto, que serão enviados
// para o banco de dados, para obter a moeda correspondente
// Se o usuário existir e tiver a moeda registrada, é retornado a moeda correspondente
// Se o usuário existir e não tiver a moeda registrada, é retornado nil
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
