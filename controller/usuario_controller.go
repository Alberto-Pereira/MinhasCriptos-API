// Package controller contém os controladores das entidades usuário e cripto
// Contém também as rotas para operações com as entidades
package controller

import (
	"minhascriptos/model"
	"minhascriptos/service"

	"github.com/gin-gonic/gin"
)

// Cadastrar Usuario
// O usuario é recebido através de uma requisição e enviado para a etapa de serviço para ser validado e cadastrado
// Se ele for cadastrado, é retornado uma resposta com número e mensagem de sucesso
// Se ele não for cadastrado, é retornado uma resposta com número e mensagem de falha
// CadastrarUsuario godoc
// @Summary Cadastra um usuário
// @Description Recebe um usuário e cadastra no sistema
// @Tags cadastrar-usuario
// @Accept json
// @Produce json
// @Success 200 {boolean} true
// @Failure 400 {object} error
// @Router /cadastrar-usuario [post]
func CadastrarUsuario(ctx *gin.Context) {

	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	isUsuarioCadastrado, status := service.CadastrarUsuario(usuario)

	if isUsuarioCadastrado == false {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

// Obter Usuario
// O usuario é recebido através de uma requisição e enviado para a etapa de serviço para ser validado e buscado
// Se ele for encontrado, é retornado uma resposta com número e o usuário buscado
// Se ele não for encontrado, é retornado uma resposta com número e mensagem de falha
func ObterUsuario(ctx *gin.Context) {

	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	u, isUsuarioValido, status := service.ObterUsuario(usuario.Email, usuario.Senha)

	if isUsuarioValido == false {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, u)
	}
}

// Obter Dinheiro Inserido
// O usuario é recebido através de uma requisição e enviado para a etapa de serviço
// para ser validado e buscar seu dinheiro inserido
// Se o dinheiro for encontrado, é retornado uma resposta com número e o dinheiro buscado
// Se o dinheiro não for encontrado, é retornado uma resposta com número e mensagem de falha
func ObterDinheiroInserido(ctx *gin.Context) {

	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	dinheiroInserido, isDinheiroInseridoValido, status := service.ObterDinheiroInserido(usuario.ID)

	if isDinheiroInseridoValido == false {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, dinheiroInserido)
	}
}
