// Package controller contém os controladores das entidades usuário e cripto
// Contém também as rotas para operações com as entidades
package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// Recebe o router
	router := gin.Default()

	// Usuário

	// Rota para cadastro de usuário
	router.POST("/usuario", CadastrarUsuario)

	// Rota para obter usuário
	router.GET("/usuario", ObterUsuario)

	// Rota para obter dinheiro inserido para aquele usuário fornecido
	router.GET("/total", ObterDinheiroInserido)

	// Cripto

	// Rota para adicionar moeda
	router.POST("/cripto", AdicionarMoeda)

	// Rota para atualizar moeda
	router.PUT("/cripto", EditarMoeda)

	// Rota para deletar moeda
	router.DELETE("/cripto", DeletarMoeda)

	// Rota para obter moedas
	router.GET("/cripto", ObterMoedas)

	// Rota para obter moedas com parâmetros personalizados
	router.GET("/criptos-busca-personalizada", ObterMoedasBuscaPersonalizada)

	return router
}
