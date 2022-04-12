package service

import (
	"fmt"
	"minhascriptos/model"
	"minhascriptos/repository"
)

func AdicionarMoeda(cripto model.Cripto) bool {
	isMoedaAdicionada := repository.AdicionarMoeda(cripto)

	if isMoedaAdicionada == 0 {
		fmt.Println("Moeda não adicionada!")
		return false
	} else {
		fmt.Println("Moeda adicionada!")
		return true
	}
}

func EditarMoeda(cripto model.Cripto) bool {
	isMoedaEditada := repository.EditarMoeda(cripto)

	if isMoedaEditada == 0 {
		fmt.Println("Moeda não editada!")
		return false
	} else {
		fmt.Println("Moeda editada!")
		return true
	}
}

func DeletarMoeda(cripto model.Cripto) bool {
	isMoedaDeletada := repository.DeletarMoeda(cripto)

	if isMoedaDeletada != cripto.ID {
		fmt.Println("Moeda não deletada! ID requisitado:", cripto.ID, " | ID encontrado:", isMoedaDeletada)
		return false
	} else {
		fmt.Println("Moeda deletada! ID:", isMoedaDeletada)
		return true
	}
}

func ObterMoedas(usuario_id int) []model.Cripto {
	moedas := repository.ObterMoedas(usuario_id)

	return moedas
}

func ObterMoedasBuscaPersonalizada(usuario_id int, tipoMoeda string, dataDeCompra string) []model.Cripto {
	moedas := repository.ObterMoedasBuscaPersonalizada(usuario_id, tipoMoeda, dataDeCompra)

	return moedas
}
