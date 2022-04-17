package service

import (
	"minhascriptos/model"
	"minhascriptos/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// Cadastrar Usuário
// Usuário válido - Deve possuir um email não cadastrado no sistema
func TestCadastrarUsuario_UsuarioValido(t *testing.T) {

	assertions := require.New(t)

	// ! Email deve ser trocado a cada teste. Não pode ser um email já cadastrado!
	usuarioValido := model.Usuario{Nome: "Teste", Email: "email12@gmail.com", Senha: "123456"}

	resultado, status := CadastrarUsuario(usuarioValido)

	assertions.Equal(resultado, true)
	assertions.Equal(status, util.HttpStatus{ID: 200, Mensagem: "Usuário cadastrado!"})

}

// Cadastrar Usuário
// Usuário inválido - A estrutura não segue as regras definidas
func TestCadastrarUsuario_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := []model.Usuario{
		{Nome: "T", Email: "email@gmail.com", Senha: "123456"}, // Nome inválido
		{Nome: "Teste", Email: "email@gmail", Senha: "123456"}, // Email inválido
		{Nome: "Teste", Email: "email@gmail.com", Senha: "12"}, // Senha inválida
	}

	for _, u := range usuarioInvalido {

		resultado, status := CadastrarUsuario(u)

		assertions.Equal(resultado, false)
		assertions.Equal(status, util.HttpStatus{ID: 400, Mensagem: "Usuário inválido!"})
	}
}

// Cadastrar Usuário
// Usuário já cadastrado - O usuário não pode ser cadastrado pois já está cadastrado
func TestCadastrarUsuario_UsuarioCadastrado(t *testing.T) {

	assertions := require.New(t)

	usuarioCadastrado := model.Usuario{Nome: "Teste", Email: "email1@gmail.com", Senha: "123456"}

	resultado, status := CadastrarUsuario(usuarioCadastrado)

	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 406, Mensagem: "Usuário já cadastrado!"})
}

// Obter Usuário
// Usuário válido - Retorna um usuário cadastrado
func TestObterUsuario_UsuarioCadastrado(t *testing.T) {

	assertions := require.New(t)

	usuarioCadastrado := model.Usuario{Email: "email1@gmail.com", Senha: "123456"}

	usuario, resultado, status := ObterUsuario(usuarioCadastrado.Email, usuarioCadastrado.Senha)

	assertions.NotEmpty(usuario)
	assertions.Equal(resultado, true)
	assertions.Equal(status, util.HttpStatus{ID: 200})
}

// Obter Usuário
// Email e/ou senha inválida - Email e/ou senha não seguem os padrões requisitados
func TestObterUsuario_EmailESenhaInvalida(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := []model.Usuario{
		{Email: "email@gmail", Senha: "123456"}, // Email inválido
		{Email: "email@gmail.com", Senha: "12"}, // Senha inválida
	}

	for _, u := range usuarioInvalido {

		u, resultado, status := ObterUsuario(u.Email, u.Senha)

		assertions.Empty(u)
		assertions.Equal(resultado, false)
		assertions.Equal(status, util.HttpStatus{ID: 400,
			Mensagem: "Email ou senha não seguem os padrões requisitados!"})

	}
}

// Obter Usuário
// Senha inválida - Email cadastrado mas senha inválida para usuário requisitado
func TestObterUsuario_SenhaInvalida(t *testing.T) {

	assertions := require.New(t)

	usuarioCadastrado := model.Usuario{Email: "email1@gmail.com", Senha: "12345"}

	usuario, resultado, status := ObterUsuario(usuarioCadastrado.Email, usuarioCadastrado.Senha)

	assertions.Empty(usuario)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 406, Mensagem: "Senha inválida para o email informado!"})
}

// Obter Usuário
// Email não cadastrado - Email não cadastrado
func TestObterUsuario_EmailNaoCadastrado(t *testing.T) {

	assertions := require.New(t)

	usuarioNaoCadastrado := model.Usuario{Email: "emailnaocadastrado@gmail.com", Senha: "123456"}

	usuario, resultado, status := ObterUsuario(usuarioNaoCadastrado.Email, usuarioNaoCadastrado.Senha)

	assertions.Empty(usuario)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Usuário não encontrado para o email informado!"})
}

// Obter dinheiro inserido
// Usuario válido - Retorna um slice com tipo e total para cada moeda
func TestObterDinheiroInserido_UsuarioValido(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{ID: 1}

	usuario, resultado, status := ObterDinheiroInserido(usuarioValido.ID)

	assertions.NotEmpty(usuario)
	assertions.Equal(resultado, true)
	assertions.Equal(status, util.HttpStatus{ID: 200})
}

// Obter dinheiro inserido
// Usuario inválido - Usuário não cadastrado
func TestObterDinheiroInserido_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := model.Usuario{ID: 999}

	usuario, resultado, status := ObterDinheiroInserido(usuarioInvalido.ID)

	assertions.Empty(usuario)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 400, Mensagem: "ID inválido!"})
}

// Obter dinheiro inserido
// Usuario válido sem dinheiro inserido - Usuário é válido mas não possui dinheiro inserido
func TestObterDinheiroInserido_UsuarioValidoSemDinheiroInserido(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := model.Usuario{ID: 12}

	usuario, resultado, status := ObterDinheiroInserido(usuarioInvalido.ID)

	assertions.Empty(usuario)
	assertions.Equal(resultado, false)
	assertions.Equal(status, util.HttpStatus{ID: 404, Mensagem: "Dinheiro inserido não encontrado para o ID informado!"})
}
