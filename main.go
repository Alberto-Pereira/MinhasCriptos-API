package main

import (
	"minhascriptos/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Usuário
	router.POST("/cadastrar-usuario", controller.CadastrarUsuario)

	router.GET("/usuario", controller.ObterUsuario)

	router.GET("/total", controller.ObterDinheiroInserido)

	// Cripto
	router.POST("/adicionar-cripto", controller.AdicionarMoeda)

	router.PUT("/editar-cripto", controller.EditarMoeda)

	router.DELETE("/deletar-cripto", controller.DeletarMoeda)

	router.GET("/criptos", controller.ObterMoedas)

	router.GET("/criptos-busca-personalizada", controller.ObterMoedasBuscaPersonalizada)

	router.Run()
}
