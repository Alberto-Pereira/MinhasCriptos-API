package service

import (
	"minhascriptos/model"
	"minhascriptos/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// Adicionar moeda
// Moeda válida - Deve estar de acordo com a estrutura e possuir um usuário vinculado
func TestAdicionarMoeda_MoedaValida(t *testing.T) {

	assertions := require.New(t)

	moedaValida := model.Cripto{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
		ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}}

	resultado, status := AdicionarMoeda(moedaValida)

	assertions.Equal(resultado, true)
	assertions.Equal(status, util.HttpStatus{ID: 200, Mensagem: "Moeda adicionada!"})
}

// Adicionar moeda
// Moeda inválida - Estrutura da moeda está inválida
func TestAdicionarMoeda_MoedaInvalida(t *testing.T) {

	assertions := require.New(t)

	moedaInvalida := []model.Cripto{
		{TipoMoeda: "bITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{TipoMoeda: "BITCOIN", DataDeCompra: "022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: -10, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: -100,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: -1000, UsuarioId: model.Usuario{ID: 11}},
		{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: -999}},
	}

	for _, mInvalida := range moedaInvalida {

		resultado, status := AdicionarMoeda(mInvalida)

		assertions.Equal(resultado, false)
		assertions.Equal(status, util.HttpStatus{ID: 400, Mensagem: "Moeda inválida!"})
	}
}

// Adicionar moeda
// Moeda válida mas usuario não cadastrado - Deve possuir uma estrutura de moeda válida mas usuário não cadastrado
func TestAdicionarMoeda_MoedaValidaMasUsuarioNaoCadastrado(t *testing.T) {

	assertions := require.New(t)

	moedaValidaMasUsuarioNaoCadastrado := model.Cripto{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17",
		QuantidadeComprada: 0.0009, PrecoDeCompra: 50, ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 999}}

	resultado, status := AdicionarMoeda(moedaValidaMasUsuarioNaoCadastrado)

	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Usuário não encontrado para o ID informado!"})
}

// Editar moeda
// Moeda válida - Deve estar de acordo com a estrutura, possuir um usuário vinculado e um id
func TestEditarMoeda_MoedaValida(t *testing.T) {

	assertions := require.New(t)

	moedaValida := model.Cripto{ID: 24, TipoMoeda: "ETH", DataDeCompra: "2022-04-20", QuantidadeComprada: 0.001,
		PrecoDeCompra: 60, ValorDaUnidadeNoDiaDeCompra: 14500, UsuarioId: model.Usuario{ID: 11}}

	resultado, status := EditarMoeda(moedaValida)

	assertions.Equal(resultado, true)
	assertions.Equal(status, util.HttpStatus{ID: 200, Mensagem: "Moeda editada!"})
}

// Editar moeda
// Moeda inválida - Estrutura da moeda está inválida
func TestEditarMoeda_MoedaInvalida(t *testing.T) {

	assertions := require.New(t)

	moedaInvalida := []model.Cripto{
		{ID: 11, TipoMoeda: "bITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{ID: 11, TipoMoeda: "BITCOIN", DataDeCompra: "022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{ID: 11, TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: -10, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{ID: 11, TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: -100,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: 11}},
		{ID: 11, TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: -1000, UsuarioId: model.Usuario{ID: 11}},
		{ID: 11, TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", QuantidadeComprada: 0.0009, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 195000, UsuarioId: model.Usuario{ID: -999}},
	}

	for _, mInvalida := range moedaInvalida {

		resultado, status := AdicionarMoeda(mInvalida)

		assertions.Equal(resultado, false)
		assertions.Equal(status, util.HttpStatus{ID: 400, Mensagem: "Moeda inválida!"})
	}
}

// Editar moeda
// Moeda válida mas id incorreto - Moeda válida mas id vinculado está incorreto
func TestEditarMoeda_MoedaValidaMasIDIncorreto(t *testing.T) {

	assertions := require.New(t)

	moedaValidaMasIDIncorreto := model.Cripto{ID: 999, TipoMoeda: "ETH", DataDeCompra: "2022-04-20", QuantidadeComprada: 0.001,
		PrecoDeCompra: 60, ValorDaUnidadeNoDiaDeCompra: 14500, UsuarioId: model.Usuario{ID: 11}}

	resultado, status := EditarMoeda(moedaValidaMasIDIncorreto)

	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Moeda não encontrada para o ID informado!"})
}

// Deletar moeda
// Moeda válida - Deve possuir apenas um id e usuário vinculado válidos
func TestDeletarMoeda_MoedaValida(t *testing.T) {

	assertions := require.New(t)

	// ! Colocar um id e usuário vinculados válidos do banco de dados
	moedaValida := model.Cripto{ID: 3, UsuarioId: model.Usuario{ID: 1}}

	resultado, status := DeletarMoeda(moedaValida)

	assertions.Equal(resultado, true)
	assertions.Equal(status, util.HttpStatus{ID: 200, Mensagem: "Moeda deletada!"})
}

// Deletar moeda
// Moeda não encontrada - O id da moeda ou do usuário vinculado está incorreto
func TestDeletarMoeda_MoedaNaoEncontrada(t *testing.T) {

	assertions := require.New(t)

	moedaNaoEncontrada := model.Cripto{ID: 999, UsuarioId: model.Usuario{ID: 999}}

	resultado, status := DeletarMoeda(moedaNaoEncontrada)

	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Moeda não encontrada para o ID informado!"})
}

// Obter moedas
// Moedas válidas - Retorna as moedas do usuário válido informado
func TestObterMoedas_MoedasValidas(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{ID: 11}

	usuario, resultado, status := ObterMoedas(usuarioValido.ID)

	assertions.NotEmpty(usuario)
	assertions.Equal(resultado, true)
	assertions.Equal(status, util.HttpStatus{ID: 200})
}

// Obter moedas
// Usuario inválido - Não consegue encontrar as moedas pois usuário é inválido
func TestObterMoedas_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := model.Usuario{ID: 999}

	usuario, resultado, status := ObterMoedas(usuarioInvalido.ID)

	assertions.Empty(usuario)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Usuário não encontrado para o ID informado!"})
}

// Obter moedas
// Moedas não cadastradas - O usuário é válido mas não possui moedas cadastradas
func TestObterMoedas_MoedasNaoCadastradas(t *testing.T) {

	assertions := require.New(t)

	usuarioValidoMoedasNaoCadastradas := model.Usuario{ID: 12}

	usuario, resultado, status := ObterMoedas(usuarioValidoMoedasNaoCadastradas.ID)

	assertions.Empty(usuario)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Moedas não encontradas para o ID informado!"})
}

// Obter moedas busca personalizada
// Usuario e parâmetros válidos - Retorna moedas para o usuário informado de acordo com os parâmetros informados
func TestObterMoedasBuscaPersonalizada_UsuarioEParametrosValidos(t *testing.T) {

	assertions := require.New(t)

	usuarioEParametrosValidos := []model.Cripto{
		{TipoMoeda: "BITCOIN", DataDeCompra: "2004-12-23", UsuarioId: model.Usuario{ID: 1}},
		{TipoMoeda: "", DataDeCompra: "2004-12-23", UsuarioId: model.Usuario{ID: 1}},
		{TipoMoeda: "BITCOIN", DataDeCompra: "", UsuarioId: model.Usuario{ID: 1}},
	}

	for _, uEParametrosValidos := range usuarioEParametrosValidos {

		moedaValida, resultado, status := ObterMoedasBuscaPersonalizada(uEParametrosValidos.UsuarioId.ID,
			uEParametrosValidos.TipoMoeda, uEParametrosValidos.DataDeCompra)

		assertions.NotEmpty(moedaValida)
		assertions.Equal(resultado, true)
		assertions.Equal(status, util.HttpStatus{ID: 200})
	}

}

// Obter moedas busca personalizada
// Parâmetros vazios - Não consegue localizar moedas pois os parâmetros estão vazios
func TestObterMoedasBuscaPersonalizada_ParametrosVazios(t *testing.T) {

	assertions := require.New(t)

	moedaComParametrosVazios := model.Cripto{TipoMoeda: "", DataDeCompra: "", UsuarioId: model.Usuario{ID: 11}}

	mComParametrosVazios, resultado, status := ObterMoedasBuscaPersonalizada(moedaComParametrosVazios.UsuarioId.ID,
		moedaComParametrosVazios.TipoMoeda, moedaComParametrosVazios.DataDeCompra)

	assertions.Empty(mComParametrosVazios)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 400, Mensagem: "O tipo de moeda e data de compra estão vazios!"})
}

// Obter moedas busca personalizada
// Tipo moeda inválido - Não consegue localizar moedas pois o tipo de moeda está incorreto
func TestObterMoedasBuscaPersonalizada_TipoMoedaInvalido(t *testing.T) {

	assertions := require.New(t)

	moedaComParametroIncorreto := model.Cripto{TipoMoeda: "bITCOIN", DataDeCompra: "", UsuarioId: model.Usuario{ID: 11}}

	mComParametroIncorreto, resultado, status := ObterMoedasBuscaPersonalizada(moedaComParametroIncorreto.UsuarioId.ID,
		moedaComParametroIncorreto.TipoMoeda, moedaComParametroIncorreto.DataDeCompra)

	assertions.Empty(mComParametroIncorreto)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 400, Mensagem: "Tipo de moeda inválido!"})
}

// Obter moedas busca personalizada
// Data de compra inválida - Não consegue localizar moedas pois a data de compra está incorreta
func TestObterMoedasBuscaPersonalizada_DataDeCompraInvalida(t *testing.T) {

	assertions := require.New(t)

	moedaComParametroIncorreto := model.Cripto{TipoMoeda: "", DataDeCompra: "022-04-17", UsuarioId: model.Usuario{ID: 11}}

	mComParametroIncorreto, resultado, status := ObterMoedasBuscaPersonalizada(moedaComParametroIncorreto.UsuarioId.ID,
		moedaComParametroIncorreto.TipoMoeda, moedaComParametroIncorreto.DataDeCompra)

	assertions.Empty(mComParametroIncorreto)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 400, Mensagem: "Data de compra inválida!"})
}

// Obter moedas busca personalizada
// Usuário inválido - Não consegue localizar moedas pois o usuário vinculado está inválido
func TestObterMoedasBuscaPersonalizada_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	moedaComParametroIncorreto := model.Cripto{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17",
		UsuarioId: model.Usuario{ID: 999}}

	mComParametroIncorreto, resultado, status := ObterMoedasBuscaPersonalizada(moedaComParametroIncorreto.UsuarioId.ID,
		moedaComParametroIncorreto.TipoMoeda, moedaComParametroIncorreto.DataDeCompra)

	assertions.Empty(mComParametroIncorreto)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Usuário não encontrado para o ID informado!"})
}

// Obter moedas busca personalizada
// Usuário sem moedas - Não consegue localizar moedas pois o usuário não possui moedas cadastradas
func TestObterMoedasBuscaPersonalizada_UsuarioSemMoedas(t *testing.T) {

	assertions := require.New(t)

	usuarioSemMoedas := model.Cripto{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17",
		UsuarioId: model.Usuario{ID: 12}}

	uSemMoedas, resultado, status := ObterMoedasBuscaPersonalizada(usuarioSemMoedas.UsuarioId.ID,
		usuarioSemMoedas.TipoMoeda, usuarioSemMoedas.DataDeCompra)

	assertions.Empty(uSemMoedas)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Moedas não encontradas para o ID informado!"})
}
