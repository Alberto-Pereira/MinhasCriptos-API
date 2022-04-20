// Package controller contém os controladores das entidades usuário e cripto
// Contém também as rotas para operações com as entidades
package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// Recebe o router
	router := gin.Default()

	// Grupo de rotas de usuário
	usuario := router.Group("/usuario")
	{
		// Rota para cadastro de usuário
		usuario.POST("/", CadastrarUsuario)

		// Rota para obter usuário
		usuario.GET("/:email/:senha", ObterUsuario)

		// Rota para obter dinheiro inserido para aquele usuário fornecido
		usuario.GET("/total/:total", ObterDinheiroInserido)
	}

	// Grupo de rotas de cripto
	cripto := router.Group("/cripto")
	{
		// Rota para adicionar moeda
		cripto.POST("/", AdicionarMoeda)

		// Rota para atualizar moeda
		cripto.PUT("/", EditarMoeda)

		// Rota para deletar moeda
		cripto.DELETE("/:idMoeda/:idUsuario", DeletarMoeda)

		// Rota para obter moedas
		cripto.GET("/:idUsuario", ObterMoedas)

		// Rota para obter moedas com parâmetros personalizados
		cripto.GET("/busca-personalizada/:idUsuario/:tipoMoeda/:dataDeCompra", ObterMoedasBuscaPersonalizada)
	}

	return router
}
