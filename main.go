package main

import (
	"minhascriptos/controller"
	"minhascriptos/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           MinhasCriptos API
// @version         1.0
// @description     API Rest para a aplicação MinhasCriptos

// @contact.name   MinhasCriptos API Suporte
// @contact.url    https://portfolio-alberto-pereira.herokuapp.com/contact
// @contact.email  alberto.pereira.dev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {

	// Recebe o router
	router := controller.SetupRouter()

	// O caminho do swagger
	docs.SwaggerInfo.BasePath = "/"

	// Rota do swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Inicializa o router
	router.Run()
}
