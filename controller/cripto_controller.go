// Package controller contém os controladores das entidades usuário e cripto
// Contém também as rotas para operações com as entidades
package controller

import (
	"minhascriptos/model"
	"minhascriptos/service"

	"github.com/gin-gonic/gin"
)

// Adicionar Moeda
// A cripto é recebida através de uma requisição e enviada para a etapa de serviço para ser validada e adicionada
// Se ela for adicionada, é retornado uma resposta com número e mensagem de sucesso
// Se ela não for adicionada, é retornado uma resposta com número e mensagem de falha
func AdicionarMoeda(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	isMoedaAdicionada, status := service.AdicionarMoeda(cripto)

	if isMoedaAdicionada == true {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

// Editar Moeda
// A cripto é recebida através de uma requisição e enviada para a etapa de serviço para ser validada e atualizada
// Se ela for atualizada, é retornado uma resposta com número e mensagem de sucesso
// Se ela não for atualizada, é retornado uma resposta com número e mensagem de falha
func EditarMoeda(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	isMoedaEditada, status := service.EditarMoeda(cripto)

	if isMoedaEditada == true {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

// Deletar Moeda
// A cripto é recebida através de uma requisição e enviada para a etapa de serviço para ser validada e deletada
// Se ela for deletada, é retornado uma resposta com número e mensagem de sucesso
// Se ela não for deletada, é retornado uma resposta com número e mensagem de falha
func DeletarMoeda(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	isMoedaDeletada, status := service.DeletarMoeda(cripto)

	if isMoedaDeletada == true {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

// Obter Moedas
// O usuario é recebido através de uma requisição e enviado para a etapa de serviço para ser validado e buscar suas moedas
// Se elas forem encontradas, é retornado uma resposta com número e as moedas associadas
// Se elas não forem encontradas, é retornado uma resposta com número e mensagem de falha
func ObterMoedas(ctx *gin.Context) {

	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	criptos, isMoedasObtidas, status := service.ObterMoedas(usuario.ID)

	if isMoedasObtidas == true {
		ctx.JSON(status.ID, criptos)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

// Obter Moedas Busca Personalizada
// A cripto é recebida através de uma requisição e enviada para a etapa de serviço para ser validada e buscar suas moedas
// usando os parâmetros personalizados definidos
// Se elas forem encontradas, é retornado uma resposta com número e as moedas associadas
// Se elas não forem encontradas, é retornado uma resposta com número e mensagem de falha
func ObterMoedasBuscaPersonalizada(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	criptos, isMoedasObtidas, status := service.ObterMoedasBuscaPersonalizada(cripto.UsuarioId.ID, cripto.TipoMoeda, cripto.DataDeCompra)

	if isMoedasObtidas == true {
		ctx.JSON(status.ID, criptos)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}
