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

	isMoedaAdicionada := service.AdicionarMoeda(cripto)

	if isMoedaAdicionada == true {
		ctx.JSON(http.StatusCreated, isMoedaAdicionada)
	} else {
		ctx.JSON(http.StatusBadRequest, isMoedaAdicionada)
	}
}

func EditarMoeda(ctx *gin.Context) {
	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isMoedaEditada := service.EditarMoeda(cripto)

	if isMoedaEditada == true {
		ctx.JSON(http.StatusAccepted, isMoedaEditada)
	} else {
		ctx.JSON(http.StatusNotAcceptable, isMoedaEditada)
	}
}

func DeletarMoeda(ctx *gin.Context) {
	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isMoedaDeletada := service.DeletarMoeda(cripto)

	if isMoedaDeletada == true {
		ctx.JSON(http.StatusAccepted, isMoedaDeletada)
	} else {
		ctx.JSON(http.StatusNotAcceptable, isMoedaDeletada)
	}
}

func ObterMoedas(ctx *gin.Context) {
	var usuario model.Usuario

	err := ctx.BindJSON(&usuario)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	criptos := service.ObterMoedas(usuario.ID)

	if criptos != nil {
		ctx.JSON(http.StatusFound, criptos)
	} else {
		ctx.JSON(http.StatusNotFound, false)
	}
}

func ObterMoedasBuscaPersonalizada(ctx *gin.Context) {
	var cripto model.Cripto

	err := ctx.BindJSON(&cripto)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	criptos := service.ObterMoedasBuscaPersonalizada(cripto.UsuarioId.ID, cripto.TipoMoeda, cripto.DataDeCompra)

	if criptos != nil {
		ctx.JSON(http.StatusFound, criptos)
	} else {
		ctx.JSON(http.StatusNotFound, false)
	}
}
