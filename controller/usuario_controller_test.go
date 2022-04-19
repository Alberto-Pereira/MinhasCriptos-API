package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Cadastrar usuário
// Usuário válido - Cadastra um usuário válido
func TestCadastrarUsuario_UsuarioValido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	// ! Email deve ser trocado a cada teste. Não pode ser um email já cadastrado!
	req, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": "Teste Controller",
		"email": "emailcontroller5@gmail.com",
		"senha": "123456"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.Equal("\"Usuário cadastrado!\"", w.Body.String())
}

// Cadastrar usuário
// Usuário incorreto - Estrutura para cadastro de usuário está incorreta
func TestCadastrarUsuario_UsuarioIncorreto(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": 123,
		"email": "emailcontroller3@gmail.com",
		"senha": "123456"
	}`))

	req2, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": "Teste Controller",
		"email": 123,
		"senha": "123456"
	}`))

	req3, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": "Teste Controller",
		"email": "emailcontroller3@gmail.com",
		"senha": 123
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())

	assertions.Equal(400, w2.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w2.Body.String())

	assertions.Equal(400, w3.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w3.Body.String())
}

// Cadastrar usuário
// Usuário inválido - Usuário não atende as regras de validação
func TestCadastrarUsuario_UsuarioInvalido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": "T",
		"email": "emailcontroller3@gmail.com",
		"senha": "123456"
	}`))

	req2, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": "Teste Controller",
		"email": "emailcontroller3@gmail",
		"senha": "123456"
	}`))

	req3, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": "Teste Controller",
		"email": "emailcontroller3@gmail.com",
		"senha": "12"
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Usuário inválido!\"", w.Body.String())

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Usuário inválido!\"", w.Body.String())

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Usuário inválido!\"", w.Body.String())
}

// Cadastrar usuário
// Usuário cadastrado - Usuário já cadastrado
func TestCadastrarUsuario_UsuarioCadastrado(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/usuario", strings.NewReader(`{
		"nome": "Teste Controller",
		"email": "emailcontroller1@gmail.com",
		"senha": "123456"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(406, w.Code)
	assertions.Equal("\"Usuário já cadastrado!\"", w.Body.String())
}

// Obter usuário
// Usuário válido - Retorna um usuário válido ao fornecer email e senha
func TestObterUsuario_UsuarioValido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/usuario", strings.NewReader(`{
		"email": "emailcontroller1@gmail.com",
		"senha": "123456"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.NotEmpty(w.Body.String())
}

// Obter usuário
// Usuário inválido - Estrutura para obter usuário está incorreta
func TestObterUsuario_UsuarioIncorreto(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/usuario", strings.NewReader(`{
		"email": 113,
		"senha": "123456"
	}`))

	req2, _ := http.NewRequest("GET", "/usuario", strings.NewReader(`{
		"email": "emailcontroller1@gmail.com",
		"senha": 123
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())
}

// Obter usuário
// Usuário incorreto - Usuário não atende as regras de validação
func TestObterUsuario_UsuarioInvalido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/usuario", strings.NewReader(`{
		"email": "emailcontroller1@gmail",
		"senha": "123456"
	}`))

	req2, _ := http.NewRequest("GET", "/usuario", strings.NewReader(`{
		"email": "emailcontroller1@gmail.com",
		"senha": "12"
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Email ou senha não seguem os padrões requisitados!\"", w.Body.String())

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Email ou senha não seguem os padrões requisitados!\"", w.Body.String())
}

// Obter usuário
// Senha inválida - A senha não corresponde ao email informado
func TestObterUsuario_UsuarioValidoSenhaInvalida(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/usuario", strings.NewReader(`{
		"email": "emailcontroller1@gmail.com",
		"senha": "12345"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(406, w.Code)
	assertions.Equal("\"Senha inválida para o email informado!\"", w.Body.String())
}

// Obter usuário
// Usuário não encontrado - Email informado não corresponde a um usuário registrado
func TestObterUsuario_UsuarioNaoEncontrado(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/usuario", strings.NewReader(`{
		"email": "emailnaoregistrado@gmail.com",
		"senha": "123456"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(404, w.Code)
	assertions.Equal("\"Usuário não encontrado para o email informado!\"", w.Body.String())
}

// Obter dinheiro inserido
// Usuário válido - Retorna o dinheiro inserido do usuário
func TestObterDinheiroInserido_UsuarioValido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/total", strings.NewReader(`{
		"user_id": 11
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(200, w.Code)
	assertions.NotEmpty(w.Body.String())
}

// Obter dinheiro inserido
// Usuário incorreto - Estrutura do usuário não está correta
func TestObterDinheiroInserido_UsuarioIncorreto(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/total", strings.NewReader(`{
		"user_id": "11"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(400, w.Code)
	assertions.Equal("\"Estrutura incorreta!\"", w.Body.String())
}

// Obter dinheiro inserido
// Usuário inválido - Usuário não registrado
func TestObterDinheiroInserido_UsuarioInvalido(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/total", strings.NewReader(`{
		"user_id": 999
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(400, w.Code)
	assertions.NotEmpty("\"ID inválido!\"", w.Body.String())
}

// Obter dinheiro inserido
// Usuário válido sem dinheiro - Usuário não possui dinheiro registrado
func TestObterDinheiroInserido_UsuarioValidoSemDinheiro(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/total", strings.NewReader(`{
		"user_id": 12
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(404, w.Code)
	assertions.NotEmpty("\"Dinheiro inserido não encontrado para o ID informado!\"", w.Body.String())
}
