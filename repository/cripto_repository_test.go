package repository

import (
	"minhascriptos/model"
	"testing"

	"github.com/stretchr/testify/require"
)

// Adicionar moeda
// Moeda válida - Registra uma moeda válida no banco de dados
func TestAdicionarMoeda_MoedaValida(t *testing.T) {

	assertions := require.New(t)

	moedaValida := model.Cripto{TipoMoeda: "DOGE", DataDeCompra: "2022-04-18", QuantidadeComprada: 50, PrecoDeCompra: 50,
		ValorDaUnidadeNoDiaDeCompra: 1, UsuarioId: model.Usuario{ID: 11}}

	id := AdicionarMoeda(moedaValida)

	assertions.NotEqual(id, 0)
}

// Adicionar moeda
// Moeda inválida - Não registra a moeda pois não existe um usuário vinculado
func TestAdicionarMoeda_MoedaInvalida(t *testing.T) {

	assertions := require.New(t)

	moedaInvalida := model.Cripto{TipoMoeda: "ETH", DataDeCompra: "2022-04-18", QuantidadeComprada: 50, PrecoDeCompra: 50,
		ValorDaUnidadeNoDiaDeCompra: 1}

	id := AdicionarMoeda(moedaInvalida)

	assertions.Equal(id, 0)
}

// Editar moeda
// Moeda válida - Edita a moeda informada
func TestEditarMoeda_MoedaValida(t *testing.T) {

	assertions := require.New(t)

	moedaValida := model.Cripto{ID: 31, TipoMoeda: "ETH", DataDeCompra: "2022-04-18", QuantidadeComprada: 50, PrecoDeCompra: 50,
		ValorDaUnidadeNoDiaDeCompra: 1, UsuarioId: model.Usuario{ID: 11}}

	id := EditarMoeda(moedaValida)

	assertions.Equal(id, moedaValida.ID)
}

// Editar moeda
// Moeda inválida - O usuário e/ou id da moeda não correspondem
func TestEditarMoeda_MoedaInvalida(t *testing.T) {

	assertions := require.New(t)

	moedaInvalida := []model.Cripto{
		{ID: 999, TipoMoeda: "ETH", DataDeCompra: "2022-04-18", QuantidadeComprada: 50, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 1, UsuarioId: model.Usuario{ID: 11}},
		{ID: 31, TipoMoeda: "ETH", DataDeCompra: "2022-04-18", QuantidadeComprada: 50, PrecoDeCompra: 50,
			ValorDaUnidadeNoDiaDeCompra: 1, UsuarioId: model.Usuario{ID: 999}},
	}

	for _, mInvalida := range moedaInvalida {

		id := EditarMoeda(mInvalida)

		assertions.Equal(id, 0)
	}
}

// Deletar moeda
// Moeda válida - Deleta a moeda informada
func TestDeletarMoeda_MoedaValida(t *testing.T) {

	assertions := require.New(t)

	// ! Informar o id e usuário de uma moeda válida a cada teste
	moedaValida := model.Cripto{ID: 31, UsuarioId: model.Usuario{ID: 11}}

	id := DeletarMoeda(moedaValida)

	assertions.Equal(id, moedaValida.ID)
}

// Deletar moeda
// Moeda inválida - O usuário e/ou id da moeda não correspondem
func TestDeletarMoeda_MoedaInvalida(t *testing.T) {

	assertions := require.New(t)

	moedaInvalida := []model.Cripto{
		{ID: 999, UsuarioId: model.Usuario{ID: 11}},
		{ID: 30, UsuarioId: model.Usuario{ID: 999}},
	}

	for _, mInvalida := range moedaInvalida {

		id := DeletarMoeda(mInvalida)

		assertions.Equal(id, 0)
	}
}

// Obter moedas
// Usuário válido - Retorna as moedas do usuário cadastrado
func TestObterMoedas_UsuarioValido(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{ID: 11}

	moedas := ObterMoedas(usuarioValido.ID)

	assertions.NotEmpty(moedas)
}

// Obter moedas
// Usuário inválido - Usuário não cadastrado
func TestObterMoedas_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := model.Usuario{ID: 999}

	moedas := ObterMoedas(usuarioInvalido.ID)

	assertions.Empty(moedas)
}

// Obter moedas
// Usuário válido mas sem moeda - Usuário cadastrado sem moedas registradas
func TestObterMoedas_UsuarioValidoMasSemMoeda(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{ID: 12}

	moedas := ObterMoedas(usuarioValido.ID)

	assertions.Empty(moedas)
}

// Obter moedas busca personalizada
// Usuário válido e 1 parâmetro - Retorna moedas para usuário cadastrado e 1 parâmetro de busca
func TestObterMoedasBuscaPersonalizada_UsuarioValidoE1Parametro(t *testing.T) {

	assertions := require.New(t)

	moedaPersonalizada := []model.Cripto{
		{TipoMoeda: "BITCOIN", UsuarioId: model.Usuario{ID: 11}},
		{DataDeCompra: "2022-04-17", UsuarioId: model.Usuario{ID: 11}},
	}

	for _, mPersonalizada := range moedaPersonalizada {

		moedas := ObterMoedasBuscaPersonalizada(mPersonalizada.UsuarioId.ID, mPersonalizada.TipoMoeda, mPersonalizada.DataDeCompra)

		assertions.NotEmpty(moedas)
	}
}

// Obter moedas busca personalizada
// Usuário válido e 2 parâmetros - Retorna moedas para usuário cadastrado e 2 parâmetros de busca
func TestObterMoedasBuscaPersonalizada_UsuarioValidoE2Parametros(t *testing.T) {

	assertions := require.New(t)

	moedaPersonalizada := model.Cripto{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", UsuarioId: model.Usuario{ID: 11}}

	moedas := ObterMoedasBuscaPersonalizada(moedaPersonalizada.UsuarioId.ID, moedaPersonalizada.TipoMoeda,
		moedaPersonalizada.DataDeCompra)

	assertions.NotEmpty(moedas)
}

// Obter moedas busca personalizada
// Usuário inválido e 1 parâmetro - Não retorna moedas pois usuário não está cadastrado e usa 1 parâmetro de busca
func TestObterMoedasBuscaPersonalizada_UsuarioInvalidoE1Parametro(t *testing.T) {

	assertions := require.New(t)

	moedaPersonalizada := []model.Cripto{
		{TipoMoeda: "BITCOIN", UsuarioId: model.Usuario{ID: 999}},
		{DataDeCompra: "2022-04-17", UsuarioId: model.Usuario{ID: 999}},
	}

	for _, mPersonalizada := range moedaPersonalizada {

		moedas := ObterMoedasBuscaPersonalizada(mPersonalizada.UsuarioId.ID, mPersonalizada.TipoMoeda, mPersonalizada.DataDeCompra)

		assertions.Empty(moedas)
	}
}

// Obter moedas busca personalizada
// Usuário inválido e 2 parâmetros - Não retorna moedas pois usuário não está cadastrado e usa 2 parâmetro de busca
func TestObterMoedasBuscaPersonalizada_UsuarioInvalidoE2Parametros(t *testing.T) {

	assertions := require.New(t)

	moedaPersonalizada := model.Cripto{TipoMoeda: "BITCOIN", DataDeCompra: "2022-04-17", UsuarioId: model.Usuario{ID: 999}}

	moedas := ObterMoedasBuscaPersonalizada(moedaPersonalizada.UsuarioId.ID, moedaPersonalizada.TipoMoeda,
		moedaPersonalizada.DataDeCompra)

	assertions.Empty(moedas)
}

// Obter moedas pelo id
// Moeda válida - Busca uma moeda pelo id e usuário associados
func TestObterMoedaPeloID_MoedaValida(t *testing.T) {

	assertions := require.New(t)

	moedaValida := model.Cripto{ID: 25, UsuarioId: model.Usuario{ID: 11}}

	id := ObterMoedaPeloID(moedaValida)

	assertions.Equal(id, moedaValida.ID)
}

// Obter moedas pelo id
// Moeda inválida - Não retorna uma moeda pelo id e usuário associados pois estão incorretos
func TestObterMoedaPeloID_MoedaInvalida(t *testing.T) {

	assertions := require.New(t)

	moedaInvalida := []model.Cripto{
		{ID: 25, UsuarioId: model.Usuario{ID: 999}},
		{ID: 999, UsuarioId: model.Usuario{ID: 11}},
	}

	for _, mInvalida := range moedaInvalida {

		id := ObterMoedaPeloID(mInvalida)

		assertions.Equal(id, 0)
	}
}
