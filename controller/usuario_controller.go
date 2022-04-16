package controller

import (
	"minhascriptos/model"
	"minhascriptos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isUsuarioCadastrado, status := service.CadastrarUsuario(usuario)

	if isUsuarioCadastrado == false {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

func ObterUsuario(ctx *gin.Context) {

	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u, isUsuarioValido, status := service.ObterUsuario(usuario.Email, usuario.Senha)

	if isUsuarioValido == false {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, u)
	}
}

func ObterDinheiroInserido(ctx *gin.Context) {

	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dinheiroInserido, isDinheiroInseridoValido, status := service.ObterDinheiroInserido(usuario.ID)

	if isDinheiroInseridoValido == false {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, dinheiroInserido)
	}
}
