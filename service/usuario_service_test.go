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
	usuarioValido := model.Usuario{Nome: "Teste", Email: "email10@gmail.com", Senha: "123456"}

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

	assertions.NotNil(usuario)
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

		assertions.NotNil(u)
		assertions.Equal(resultado, false)
		assertions.Equal(status, util.HttpStatus{ID: 400,
			Mensagem: "Email ou senha não seguem os padrões requisitados!"})

	}
}

// Obter Usuário
// Senha inválida - Senha inválida para usuário requisitado
func TestObterUsuario_SenhaInvalida(t *testing.T) {

}
