// Package controller contém os controladores das entidades usuário e cripto
// Contém também as rotas para operações com as entidades
package controller

import (
	"fmt"
	"minhascriptos/model"
	"minhascriptos/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Adicionar Moeda
// A cripto é recebida através de uma requisição e enviada para a etapa de serviço para ser validada e adicionada
// Se ela for adicionada, é retornado uma resposta com número e mensagem de sucesso
// Se ela não for adicionada, é retornado uma resposta com número e mensagem de falha
// AdicionarMoeda godoc
// @Summary Adiciona uma moeda
// @Description Retorna uma mensagem associada a operação
// @Tags Cripto
// @Accept application/json
// @Produce application/json
// @Param cripto body model.Cripto false "Informar o tipo de moeda(UPPERCASE), data de compra(YYYY/MM/DD), quantidade comprada, preço de compra, valor da unidade no dia de compra e o id do usuário associado."
// @Success 200 {string} mensagem
// @Failure 400 {string} mensagem
// @Failure 404 {string} mensagem
// @Failure 500 {string} mensagem
// @Router /cripto/ [post]
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
// EditarMoeda godoc
// @Summary Atualiza uma moeda
// @Description Retorna uma mensagem associada a operação
// @Tags Cripto
// @Accept application/json
// @Produce application/json
// @Param cripto body model.Cripto false "Informar o id da moeda, tipo de moeda(UPPERCASE), data de compra(YYYY/MM/DD), quantidade comprada, preço de compra, valor da unidade no dia de compra e o id do usuário associado."
// @Success 200 {string} mensagem
// @Failure 400 {string} mensagem
// @Failure 404 {string} mensagem
// @Failure 500 {string} mensagem
// @Router /cripto/ [put]
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
// DeletarMoeda godoc
// @Summary Deleta uma moeda
// @Description Retorna uma mensagem associada a operação
// @Tags Cripto
// @Accept application/json
// @Produce application/json
// @Param idMoeda path int true "Informar o id da moeda"
// @Param idUsuario path int true "Informar o id do usuário"
// @Success 200 {string} mensagem
// @Failure 400 {string} mensagem
// @Failure 404 {string} mensagem
// @Failure 500 {string} mensagem
// @Router /cripto/{idMoeda}/{idUsuario} [delete]
func DeletarMoeda(ctx *gin.Context) {

	idMoeda := ctx.Param("idMoeda")
	idUsuario := ctx.Param("idUsuario")

	if idMoeda == "" || idUsuario == "" {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	intIdMoeda, errMoeda := strconv.Atoi(idMoeda)
	intIdUsuario, errUsuario := strconv.Atoi(idUsuario)

	if errMoeda != nil || errUsuario != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	cripto := model.Cripto{ID: intIdMoeda, UsuarioId: model.Usuario{ID: intIdUsuario}}

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
// ObterMoedas godoc
// @Summary Obtém moedas de um usuário
// @Description Retorna uma mensagem associada a operação
// @Tags Cripto
// @Accept application/json
// @Produce application/json
// @Param idUsuario path int true "Informar o id do usuário que deseja obter moedas"
// @Success 200 {string} []model.Cripto
// @Failure 400 {string} mensagem
// @Failure 404 {string} mensagem
// @Router /cripto/{idUsuario} [get]
func ObterMoedas(ctx *gin.Context) {

	idUsuario := ctx.Param("idUsuario")

	if idUsuario == "" {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	intIdUsuario, err := strconv.Atoi(idUsuario)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	usuario := model.Usuario{ID: intIdUsuario}

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
// ObterMoedasBuscaPersonalizada godoc
// @Summary Obtém moedas de um usuário com parâmetros personalizados
// @Description Retorna uma mensagem associada a operação
// @Tags Cripto
// @Accept application/json
// @Produce application/json
// @Param idUsuario path int true "Informar o id do usuário que deseja obter moedas"
// @Param tipoMoeda path string false "Informar o parâmetro tipo de moeda (UPPERCASE)"
// @Param dataDeCompra path string false "Informar o parâmetro data de compra (YYYY-MM-DD)"
// @Success 200 {string} []model.Cripto
// @Failure 400 {string} mensagem
// @Failure 404 {string} mensagem
// @Router /cripto/busca-personalizada/{idUsuario}/{tipoMoeda}/{dataDeCompra} [get]
func ObterMoedasBuscaPersonalizada(ctx *gin.Context) {

	idUsuario := ctx.Param("idUsuario")
	tipoMoeda := ctx.Param("tipoMoeda")
	dataDeCompra := ctx.Param("dataDeCompra")

	fmt.Println(tipoMoeda)

	if idUsuario == "" {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	if tipoMoeda == "" && dataDeCompra == "" {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	intIdUsuario, err := strconv.Atoi(idUsuario)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	if tipoMoeda == "{tipoMoeda}" || tipoMoeda == "undefined" {
		tipoMoeda = ""
	}

	if dataDeCompra == "{dataDeCompra}" || dataDeCompra == "undefined" {
		dataDeCompra = ""
	}

	cripto := model.Cripto{TipoMoeda: tipoMoeda, DataDeCompra: dataDeCompra, UsuarioId: model.Usuario{ID: intIdUsuario}}

	criptos, isMoedasObtidas, status := service.ObterMoedasBuscaPersonalizada(cripto.UsuarioId.ID, cripto.TipoMoeda, cripto.DataDeCompra)

	if isMoedasObtidas == true {
		ctx.JSON(status.ID, criptos)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}
