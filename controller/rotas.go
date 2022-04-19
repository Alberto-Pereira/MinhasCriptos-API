package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	// Usuário
	router.POST("/usuario", CadastrarUsuario)

	router.GET("/usuario", ObterUsuario)

	router.GET("/total", ObterDinheiroInserido)

	// Cripto
	router.POST("/cripto", AdicionarMoeda)

	router.PUT("/cripto", EditarMoeda)

	router.DELETE("/cripto", DeletarMoeda)

	router.GET("/cripto", ObterMoedas)

	router.GET("/criptos-busca-personalizada", ObterMoedasBuscaPersonalizada)

	return router
}
