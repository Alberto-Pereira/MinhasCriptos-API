package service

import (
	"minhascriptos/model"
	"minhascriptos/repository"
	"minhascriptos/util"
	"regexp"
)

func AdicionarMoeda(cripto model.Cripto) (bool, util.HttpStatus) {

	if isMoedaValida(cripto) == false {
		return false, util.HttpStatus{ID: 400, Mensagem: "Moeda inválida!"}
	}

	if isUsuarioIDValido(cripto.UsuarioId.ID) == false {
		return false, util.HttpStatus{ID: 404, Mensagem: "Usuário não encontrado para o ID informado!"}
	}

	isMoedaAdicionadaID := repository.AdicionarMoeda(cripto)

	if isMoedaAdicionadaID == 0 {
		return false, util.HttpStatus{ID: 500, Mensagem: "Erro desconhecido ao tentar adicionar moeda!"}
	} else {
		return true, util.HttpStatus{ID: 200, Mensagem: "Moeda adicionada!"}
	}
}

func EditarMoeda(cripto model.Cripto) (bool, util.HttpStatus) {

	if isMoedaValida(cripto) == false {
		return false, util.HttpStatus{ID: 400, Mensagem: "Moeda inválida!"}
	}

	if isMoedaEUsuarioIDValido(cripto) == false {
		return false, util.HttpStatus{ID: 404, Mensagem: "Moeda não encontrada para o ID informado!"}
	}

	isMoedaEditadaID := repository.EditarMoeda(cripto)

	if isMoedaEditadaID != cripto.ID {
		return false, util.HttpStatus{ID: 500, Mensagem: "Erro desconhecido ao tentar editar moeda!"}
	} else {
		return true, util.HttpStatus{ID: 200, Mensagem: "Moeda editada!"}
	}
}

func DeletarMoeda(cripto model.Cripto) (bool, util.HttpStatus) {

	if isMoedaEUsuarioIDValido(cripto) == false {
		return false, util.HttpStatus{ID: 404, Mensagem: "Moeda não encontrada para o ID informado!"}
	}

	isMoedaDeletada := repository.DeletarMoeda(cripto)

	if isMoedaDeletada != cripto.ID {
		return false, util.HttpStatus{ID: 500, Mensagem: "Erro desconhecido ao deletar moeda!"}
	} else {
		return true, util.HttpStatus{ID: 200, Mensagem: "Moeda deletada!"}
	}
}

func ObterMoedas(usuario_id int) ([]model.Cripto, bool, util.HttpStatus) {

	if isUsuarioIDValido(usuario_id) == false {
		return []model.Cripto{}, false, util.HttpStatus{ID: 404, Mensagem: "Usuário não encontrado para o ID informado!"}
	}

	moedas := repository.ObterMoedas(usuario_id)

	if moedas == nil {
		return []model.Cripto{}, false, util.HttpStatus{ID: 404, Mensagem: "Moedas não encontradas para o ID informado!"}
	} else {
		return moedas, true, util.HttpStatus{ID: 200}
	}
}

func ObterMoedasBuscaPersonalizada(usuario_id int, tipoMoeda string, dataDeCompra string) ([]model.Cripto, bool, util.HttpStatus) {

	if tipoMoeda == "" && dataDeCompra == "" {
		return []model.Cripto{}, false, util.HttpStatus{ID: 400,
			Mensagem: "O tipo de moeda e data de compra estão vazios!"}
	}

	if tipoMoeda != "" {
		if isTipoMoedaValido(tipoMoeda) == false {
			return []model.Cripto{}, false, util.HttpStatus{ID: 400, Mensagem: "Tipo de moeda inválido!"}
		}
	}

	if dataDeCompra != "" {
		if isDataDeCompraValida(dataDeCompra) == false {
			return []model.Cripto{}, false, util.HttpStatus{ID: 400, Mensagem: "Data de compra inválida!"}
		}
	}

	if isUsuarioIDValido(usuario_id) == false {
		return []model.Cripto{}, false, util.HttpStatus{ID: 404, Mensagem: "Usuário não encontrado para o ID informado!"}
	}

	moedas := repository.ObterMoedasBuscaPersonalizada(usuario_id, tipoMoeda, dataDeCompra)

	if moedas == nil {
		return []model.Cripto{}, false, util.HttpStatus{ID: 404, Mensagem: "Moedas não encontradas para o ID informado!"}
	} else {
		return moedas, true, util.HttpStatus{ID: 200}
	}
}

func isMoedaValida(cripto model.Cripto) bool {

	if isTipoMoedaValido(cripto.TipoMoeda) {
		if isDataDeCompraValida(cripto.DataDeCompra) {
			if cripto.QuantidadeComprada > 0 {
				if cripto.PrecoDeCompra > 0 {
					if cripto.ValorDaUnidadeNoDiaDeCompra > 0 {
						if cripto.UsuarioId.ID > 0 {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

func isMoedaEUsuarioIDValido(cripto model.Cripto) bool {

	id := repository.ObterMoedaPeloID(cripto)

	if id != cripto.ID {
		return false
	} else {
		return true
	}
}

func isTipoMoedaValido(tipoMoeda string) bool {

	tipoMoedaValido := regexp.MustCompile(`^[A-Z]+$`)

	if tipoMoedaValido.MatchString(tipoMoeda) {
		return true
	}

	return false
}

func isDataDeCompraValida(dataDeCompra string) bool {

	dataDeCompraValida := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	if dataDeCompraValida.MatchString(dataDeCompra) {
		return true
	}

	return false
}
