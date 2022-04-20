// Package controller contém os controladores das entidades usuário e cripto
// Contém também as rotas para operações com as entidades
package controller

import (
	"minhascriptos/model"
	"minhascriptos/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Cadastrar Usuario
// O usuario é recebido através de uma requisição e enviado para a etapa de serviço para ser validado e cadastrado
// Se ele for cadastrado, é retornado uma resposta com número e mensagem de sucesso
// Se ele não for cadastrado, é retornado uma resposta com número e mensagem de falha
// CadastrarUsuario godoc
// @Summary Cadastra um usuário
// @Description Retorna uma mensagem associada a operação
// @Tags Usuário
// @Accept application/json
// @Produce application/json
// @Param usuario body model.Usuario false "Informar nome, email e senha."
// @Success 200 {string} mensagem
// @Failure 400 {string} mensagem
// @Failure 406 {string} mensagem
// @Failure 500 {string} mensagem
// @Router /usuario/ [post]
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
// ObterUsuario godoc
// @Summary Obtém um usuário
// @Description Retorna um usuário cadastrado
// @Tags Usuário
// @Accept application/json
// @Produce application/json
// @Param email path string true "Informar o email do usuário a ser encontrado."
// @Param senha path string true "Informar a senha correspondente ao email."
// @Success 200 {string} model.Usuario
// @Failure 400 {string} util.HttpStatus.Mensagem
// @Failure 404 {string} util.HttpStatus.Mensagem
// @Failure 406 {string} util.HttpStatus.Mensagem
// @Failure 500 {string} util.HttpStatus.Mensagem
// @Router /usuario/{email}/{senha} [get]
func ObterUsuario(ctx *gin.Context) {

	email := ctx.Param("email")
	senha := ctx.Param("senha")

	if email == "" || senha == "" {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	usuario := model.Usuario{Email: email, Senha: senha}

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
// ObterDinheiroInserido godoc
// @Summary Obtém o dinheiro inserido de um usuário cadastrado
// @Description Retorna o dinheiro de um usuário cadastrado
// @Tags Usuário
// @Accept application/json
// @Produce application/json
// @Param total path int true "Informar o id do usuário a qual desejar buscar o dinheiro inserido."
// @Success 200 {string} []model.DinheiroInserido
// @Failure 400 {string} mensagem
// @Failure 404 {string} mensagem
// @Router /usuario/total/{total} [get]
func ObterDinheiroInserido(ctx *gin.Context) {

	total := ctx.Param("total")

	if total == "" {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	intTotal, err := strconv.Atoi(total)

	if err != nil {
		ctx.JSON(400, "Estrutura incorreta!")
		return
	}

	usuario := model.Usuario{ID: intTotal}

	dinheiroInserido, isDinheiroInseridoValido, status := service.ObterDinheiroInserido(usuario.ID)

	if isDinheiroInseridoValido == false {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, dinheiroInserido)
	}
}
