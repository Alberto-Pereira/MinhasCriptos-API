package main

import (
	"minhascriptos/controller"

	"github.com/gin-gonic/gin"
)

// @title           MinhasCriptos API
// @version         1.0
// @description     REST API para a aplicação MinhasCriptos

// @contact.name   MinhasCriptos API Suporte
// @contact.url    https://portfolio-alberto-pereira.herokuapp.com/contact
// @contact.email  alberto.pereira.dev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {

	router := gin.Default()

	// Usuário
	router.POST("/usuario", controller.CadastrarUsuario)

	router.GET("/usuario", controller.ObterUsuario)

	router.GET("/total", controller.ObterDinheiroInserido)

	// Cripto
	router.POST("/cripto", controller.AdicionarMoeda)

	router.PUT("/cripto", controller.EditarMoeda)

	router.DELETE("/cripto", controller.DeletarMoeda)

	router.GET("/cripto", controller.ObterMoedas)

	router.GET("/criptos-busca-personalizada", controller.ObterMoedasBuscaPersonalizada)

	router.Run()
}
