package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Adicionar moeda
// Moeda válida - Adiciona uma moeda válida com usuário válido
func TestAdicionarMoeda_MoedaValida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.Equal("\"Moeda adicionada!\"", w.Body.String())
}

// Adicionar moeda
// Moeda incorreta - Moeda não está de acordo com a estrutura solicitada
func TestAdicionarMoeda_MoedaIncorreta(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	w4 := httptest.NewRecorder()
	w5 := httptest.NewRecorder()
	w6 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": 123,
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req2, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": 123,
		"quantidade_comprada": -0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req3, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": "0.001",
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req4, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": "12.50",
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req5, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": "14500",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req6, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": "11"
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)
	router.ServeHTTP(w4, req4)
	router.ServeHTTP(w5, req5)
	router.ServeHTTP(w6, req6)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())

	assertions.Equal(400, w2.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w2.Body.String())

	assertions.Equal(400, w3.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w3.Body.String())

	assertions.Equal(400, w4.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w4.Body.String())

	assertions.Equal(400, w5.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w5.Body.String())

	assertions.Equal(400, w6.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w6.Body.String())
}

// Adicionar moeda
// Moeda inválida - Moeda não está de acordo com as regras solicitadas
func TestAdicionarMoeda_MoedaInvalida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	w4 := httptest.NewRecorder()
	w5 := httptest.NewRecorder()
	w6 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "eTH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req2, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req3, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": -0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req4, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": -12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req5, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": -14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req6, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": -11
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)
	router.ServeHTTP(w4, req4)
	router.ServeHTTP(w5, req5)
	router.ServeHTTP(w6, req6)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Moeda inválida!\"", w.Body.String())

	assertions.Equal(400, w2.Code)
	assertions.Equal("\"Moeda inválida!\"", w2.Body.String())

	assertions.Equal(400, w3.Code)
	assertions.Equal("\"Moeda inválida!\"", w3.Body.String())

	assertions.Equal(400, w4.Code)
	assertions.Equal("\"Moeda inválida!\"", w4.Body.String())

	assertions.Equal(400, w5.Code)
	assertions.Equal("\"Moeda inválida!\"", w5.Body.String())

	assertions.Equal(400, w6.Code)
	assertions.Equal("\"Moeda inválida!\"", w6.Body.String())
}

// Adicionar moeda
// Moeda válida usuário inválido - Não adiciona moeda pois não encontra usuário associado
func TestAdicionarMoeda_MoedaValidaUsuarioInvalido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/cripto", strings.NewReader(`{
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 999
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(404, w.Code)
	assertions.Equal("\"Usuário não encontrado para o ID informado!\"", w.Body.String())
}

// Editar moeda
// Moeda válida - Edita uma moeda válida com usuário válido
func TestEditarMoeda_MoedaValida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 25,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.Equal("\"Moeda editada!\"", w.Body.String())
}

// Editar moeda
// Moeda incorreta - Moeda não está de acordo com a estrutura solicitada
func TestEditarMoeda_MoedaIncorreta(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	w4 := httptest.NewRecorder()
	w5 := httptest.NewRecorder()
	w6 := httptest.NewRecorder()
	w7 := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": 123,
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req2, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": 123,
		"quantidade_comprada": -0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req3, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": "0.001",
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req4, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": "12.50",
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req5, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": "14500",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req6, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": "11"
		}
	}`))

	req7, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": "1",
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)
	router.ServeHTTP(w4, req4)
	router.ServeHTTP(w5, req5)
	router.ServeHTTP(w6, req6)
	router.ServeHTTP(w7, req7)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())

	assertions.Equal(400, w2.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w2.Body.String())

	assertions.Equal(400, w3.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w3.Body.String())

	assertions.Equal(400, w4.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w4.Body.String())

	assertions.Equal(400, w5.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w5.Body.String())

	assertions.Equal(400, w6.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w6.Body.String())

	assertions.Equal(400, w7.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w7.Body.String())
}

// Editar moeda
// Moeda inválida - Moeda não está de acordo com as regras solicitadas
func TestEditarMoeda_MoedaInvalida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	w4 := httptest.NewRecorder()
	w5 := httptest.NewRecorder()
	w6 := httptest.NewRecorder()
	w7 := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": -999
		}
	}`))

	req2, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "eTH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req3, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req4, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": -0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req5, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": -12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req6, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": -14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req7, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 1,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": -11
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)
	router.ServeHTTP(w4, req4)
	router.ServeHTTP(w5, req5)
	router.ServeHTTP(w6, req6)
	router.ServeHTTP(w7, req7)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Moeda inválida!\"", w.Body.String())

	assertions.Equal(400, w2.Code)
	assertions.Equal("\"Moeda inválida!\"", w2.Body.String())

	assertions.Equal(400, w3.Code)
	assertions.Equal("\"Moeda inválida!\"", w3.Body.String())

	assertions.Equal(400, w4.Code)
	assertions.Equal("\"Moeda inválida!\"", w4.Body.String())

	assertions.Equal(400, w5.Code)
	assertions.Equal("\"Moeda inválida!\"", w5.Body.String())

	assertions.Equal(400, w6.Code)
	assertions.Equal("\"Moeda inválida!\"", w6.Body.String())

	assertions.Equal(400, w7.Code)
	assertions.Equal("\"Moeda inválida!\"", w7.Body.String())
}

// Editar moeda
// Moeda e usuário não iguais - Não encontra a moeda pois o id e o usuário informados não são iguais
func TestEditarMoeda_MoedaEUsuarioNaoIguais(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 25,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 999
		}
	}`))

	req2, _ := http.NewRequest("PUT", "/cripto", strings.NewReader(`{
		"cripto_id": 999,
		"tipo_moeda": "ETH",
		"data_de_compra": "2022-04-18",
		"quantidade_comprada": 0.001,
		"preco_de_compra": 12.50,
		"valor_da_unidade_no_dia_de_compra": 14500,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(404, w.Code)
	assertions.Equal("\"Moeda não encontrada para o ID informado!\"", w.Body.String())

	assertions.Equal(404, w2.Code)
	assertions.Equal("\"Moeda não encontrada para o ID informado!\"", w2.Body.String())
}

// Deletar moeda
// Moeda válida - Deleta uma moeda válida com usuário válido
func TestDeletarMoeda_MoedaValida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	// ! Colocar uma moeda com id e usuário válidos a cada teste
	req, _ := http.NewRequest("DELETE", "/cripto", strings.NewReader(`{
		"cripto_id": 6,
		"usuario_id": {
			"user_id": 1
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.Equal("\"Moeda deletada!\"", w.Body.String())
}

// Deletar moeda
// Moeda incorreta - Moeda não está de acordo com a estrutura solicitada
func TestDeletarMoeda_MoedaIncorreta(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/cripto", strings.NewReader(`{
		"cripto_id": "5",
		"usuario_id": {
			"user_id": 1
		}
	}`))

	req2, _ := http.NewRequest("DELETE", "/cripto", strings.NewReader(`{
		"cripto_id": 5,
		"usuario_id": {
			"user_id": "1"
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())
}

// Deletar moeda
// Moeda inválida - Moeda informada não possui usuário correto
func TestDeletarMoeda_MoedaInvalida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/cripto", strings.NewReader(`{
		"cripto_id": 999,
		"usuario_id": {
			"user_id": 1
		}
	}`))

	req2, _ := http.NewRequest("DELETE", "/cripto", strings.NewReader(`{
		"cripto_id": 5,
		"usuario_id": {
			"user_id": 999
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(404, w.Code)
	assertions.Equal("\"Moeda não encontrada para o ID informado!\"", w.Body.String())

	assertions.Equal(404, w.Code)
	assertions.Equal("\"Moeda não encontrada para o ID informado!\"", w.Body.String())
}

// Obter moeda
// Usuário válido - Retorna as moedas do usuário informado
func TestObterMoedas_UsuarioValido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/cripto", strings.NewReader(`{
		"user_id": 11
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.NotEmpty(w.Body.String())
}

// Obter moeda
// Usuário incorreto - Usuário não corresponde a estrutura solicitada
func TestObterMoedas_UsuarioIncorreto(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/cripto", strings.NewReader(`{
		"user_id": "11"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())
}

// Obter moeda
// Usuário inválido - Usuário não encontrado para id informado
func TestObterMoedas_UsuarioInvalido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/cripto", strings.NewReader(`{
		"user_id": 999
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(404, w.Code)
	assertions.Equal("\"Usuário não encontrado para o ID informado!\"", w.Body.String())
}

// Obter moeda
// Usuário válido sem moedas - Usuário é válido mas não possui moedas registradas
func TestObterMoedas_UsuarioValidoSemMoedas(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/cripto", strings.NewReader(`{
		"user_id": 12
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(404, w.Code)
	assertions.Equal("\"Moedas não encontradas para o ID informado!\"", w.Body.String())
}

// Obter moeda busca personalizada
// Usuário válido 1 parâmetro - Usuário é válido e possui 1 parâmetro de busca
func TestObterMoedasBuscaPersonalizada_UsuarioValido1Parâmetro(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"ETH",
		"data_de_compra":"",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req2, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"",
		"data_de_compra":"2022-04-18",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(200, w.Code)
	assertions.NotEmpty(w.Body.String())

	assertions.Equal(200, w2.Code)
	assertions.NotEmpty(w2.Body.String())
}

// Obter moeda busca personalizada
// Usuário válido 2 parâmetros - Usuário é válido e possui 2 parâmetros de busca
func TestObterMoedasBuscaPersonalizada_UsuarioValido2Parâmetro(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"ETH",
		"data_de_compra":"2022-04-18",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.NotEmpty(w.Body.String())
}

// Obter moeda busca personalizada
// Usuário incorreto - Usuário não segue a estrutura solicitada
func TestObterMoedasBuscaPersonalizada_UsuarioIncorreto(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda": 123,
		"data_de_compra":"",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req2, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"",
		"data_de_compra":123,
		"usuario_id": {
			"user_id": 11
		}
	}`))

	req3, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"ETH",
		"data_de_compra":"2022-04-18",
		"usuario_id": {
			"user_id": "11"
		}
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)

	assertions.Equal(400, w.Code)
	assertions.NotEmpty("\"Estrutura incorreta!\"", w.Body.String())

	assertions.Equal(400, w2.Code)
	assertions.NotEmpty("\"Estrutura incorreta!\"", w2.Body.String())

	assertions.Equal(400, w3.Code)
	assertions.NotEmpty("\"Estrutura incorreta!\"", w3.Body.String())
}

// Obter moeda busca personalizada
// Usuário vazio - O usuário se encontra vazio
func TestObterMoedasBuscaPersonalizada_UsuarioVazio(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"",
		"data_de_compra":"",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(400, w.Code)
	assertions.NotEmpty("\"O tipo de moeda e data de compra estão vazios!\"", w.Body.String())
}

// Obter moeda busca personalizada
// Tipo de moeda inválido - O tipo de moeda informado é inválido para as regras definidas
func TestObterMoedasBuscaPersonalizada_TipoDeMoedaInvalido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"eTH",
		"data_de_compra":"",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(400, w.Code)
	assertions.NotEmpty("\"Tipo de moeda inválido!\"", w.Body.String())
}

// Obter moeda busca personalizada
// Data de compra inválida - A data de compra informada é inválida para as regras definidas
func TestObterMoedasBuscaPersonalizada_DataDeCompraInvalida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"ETH",
		"data_de_compra":"022-04-18",
		"usuario_id": {
			"user_id": 11
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(400, w.Code)
	assertions.NotEmpty("\"Data de compra inválida!\"", w.Body.String())
}

// Obter moeda busca personalizada
// Usuário inválido - O usuário informado não existe
func TestObterMoedasBuscaPersonalizada_UsuarioInvalido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"ETH",
		"data_de_compra":"2022-04-18",
		"usuario_id": {
			"user_id": 999
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(404, w.Code)
	assertions.NotEmpty("\"Usuário não encontrado para o ID informado!\"", w.Body.String())
}

// Obter moeda busca personalizada
// Usuário válido sem moedas - O usuário informado existe mas não possui moedas registradas
func TestObterMoedasBuscaPersonalizada_UsuarioValidoSemMoedas(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/criptos-busca-personalizada", strings.NewReader(`{
		"tipo_moeda":"ETH",
		"data_de_compra":"2022-04-18",
		"usuario_id": {
			"user_id": 12
		}
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(404, w.Code)
	assertions.NotEmpty("\"Moedas não encontradas para o ID informado!\"", w.Body.String())
}
