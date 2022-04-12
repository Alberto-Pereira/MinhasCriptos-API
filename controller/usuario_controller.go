package controller

import (
	"minhascriptos/model"
	"minhascriptos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CadastrarUsuario(ctx *gin.Context) {
	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u := service.CadastrarUsuario(usuario, usuario.Email)

	if u == true {
		ctx.JSON(http.StatusAccepted, u)
	} else {
		ctx.JSON(http.StatusConflict, u)
	}

}

func ObterUsuario(ctx *gin.Context) {
	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u, isUsuarioValido := service.ObterUsuario(usuario.Email, usuario.Senha)

	if isUsuarioValido == false {
		ctx.JSON(http.StatusNotFound, u)
	} else {
		ctx.JSON(http.StatusAccepted, u)
	}

}

func ObterDinheiroInserido(ctx *gin.Context) {
	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dinheiroInserido := service.ObterDinheiroInserido(usuario.ID)

	ctx.JSON(http.StatusFound, dinheiroInserido)
}
