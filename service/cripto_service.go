// Package service contém as regras de serviço das entidades usuário e cripto
package service

import (
	"minhascriptos/model"
	"minhascriptos/repository"
	"minhascriptos/util"
	"regexp"
)

// Adicionar Moeda
// A entidade cripto, com usuário associado, após ser validada, é enviada para o repositório para ser adicionada
// Se ela for adicionada, é retornado true e o status associado
// Se ela não for adicionada, é retornado false e o status associado
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

// Editar Moeda
// A entidade cripto, com id e usuário associado, após ser validada, é enviada para o repositório para ser atualizada
// Se ela for atualizada, é retornado true e o status associado
// Se ela não for atualizada, é retornado false e o status associado
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

// Deletar Moeda
// A entidade cripto, com id e usuário associado, após ser validada, é enviada para o repositório para ser deletada
// Se ela for deletada, é retornado true e o status associado
// Se ela não for deletada, é retornado false e o status associado
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

// Obter Moedas
// A entidade usuario, após ser validada, é enviada para o repositório para obter suas moedas
// Se elas forem obtidas, é retornado as moedas associadas, true e o status associado
// Se elas não forem obtidas, é retornado nil, false e o status associado
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

// Obter Moedas Busca Personalizada
// A entidade usuario e utilizando parâmetros de busca,
// após ser validada, é enviada para o repositório para buscar suas moedas
// Se elas forem obtidas, é retornado as moedas associadas, true e o status associado
// Se elas não forem obtidas, é retornado nil, false e o status associado
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

// Is Moeda Valida
// Recebe uma entidade cripto para ser validada
// Se ela atender todas as regras definidas, é retornado true
// Se ela não atender uma das regras definidas, é retornado false
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

// Is Moeda E Usuario ID Valido
// Recebe uma entidade cripto com usuário relacionado para ser buscada
// Se ela existir, é retornado true
// Se ela não existir, é retornado false
func isMoedaEUsuarioIDValido(cripto model.Cripto) bool {

	id := repository.ObterMoedaPeloID(cripto)

	if id != cripto.ID {
		return false
	} else {
		return true
	}
}

// Is Tipo Moeda Valido
// Recebe um tipo de moeda para ser validado
// Se ela atender as regras definidas, é retornado true
// Se ela não atender as regras definidas, é retornado false
func isTipoMoedaValido(tipoMoeda string) bool {

	// ! Formato aceito - Deve conter todas letras maiúsculas e pelo menos uma letra
	tipoMoedaValido := regexp.MustCompile(`^[A-Z]+$`)

	if tipoMoedaValido.MatchString(tipoMoeda) {
		return true
	}

	return false
}

// Is Tipo Data De Compra Valida
// Recebe uma data de compra para ser validada
// Se ela atender as regras definidas, é retornado true
// Se ela não atender as regras definidas, é retornado false
func isDataDeCompraValida(dataDeCompra string) bool {

	// ! Formato aceito - YYYY/MM/DD
	dataDeCompraValida := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	if dataDeCompraValida.MatchString(dataDeCompra) {
		return true
	}

	return false
}
