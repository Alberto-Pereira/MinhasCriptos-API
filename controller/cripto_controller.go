package controller

import (
	"minhascriptos/model"
	"minhascriptos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdicionarMoeda(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isMoedaAdicionada, status := service.AdicionarMoeda(cripto)

	if isMoedaAdicionada == true {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

func EditarMoeda(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isMoedaEditada, status := service.EditarMoeda(cripto)

	if isMoedaEditada == true {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

func DeletarMoeda(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isMoedaDeletada, status := service.DeletarMoeda(cripto)

	if isMoedaDeletada == true {
		ctx.JSON(status.ID, status.Mensagem)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

func ObterMoedas(ctx *gin.Context) {

	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	criptos, isMoedasObtidas, status := service.ObterMoedas(usuario.ID)

	if isMoedasObtidas == true {
		ctx.JSON(status.ID, criptos)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}

func ObterMoedasBuscaPersonalizada(ctx *gin.Context) {

	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	criptos, isMoedasObtidas, status := service.ObterMoedasBuscaPersonalizada(cripto.UsuarioId.ID, cripto.TipoMoeda, cripto.DataDeCompra)

	if isMoedasObtidas == true {
		ctx.JSON(status.ID, criptos)
	} else {
		ctx.JSON(status.ID, status.Mensagem)
	}
}
