package repository

import (
	"minhascriptos/model"
	"testing"

	"github.com/stretchr/testify/require"
)

// Autenticar usuário
// Usuário cadastrado - O usuário está cadastrado para o email informado
func TestAutenticarUsuario_UsuarioCadastrado(t *testing.T) {

	assertions := require.New(t)

	usuarioCadastrado := model.Usuario{Email: "email1@gmail.com"}

	usuario := AutenticarUsuario(usuarioCadastrado.Email)

	assertions.NotEmpty(usuario)
}

// Autenticar usuário
// Usuário não cadastrado - O usuário não está cadastrado para o email informado
func TestAutenticarUsuario_UsuarioNaoCadastrado(t *testing.T) {

	assertions := require.New(t)

	usuarioNaoCadastrado := model.Usuario{Email: "emailnaocadastrado@gmail.com"}

	usuario := AutenticarUsuario(usuarioNaoCadastrado.Email)

	assertions.Empty(usuario)
}

// Cadastrar usuário
// Usuário válido - Cadastra um usuário válido no banco de dados
func TestCadastrarUsuario_UsuarioValido(t *testing.T) {

	assertions := require.New(t)

	// ! Email deve ser trocado a cada teste. Não pode ser um email já cadastrado!
	usuarioValido := model.Usuario{Nome: "Teste Repository", Email: "email18@gmail.com", Senha: "123456"}

	usuario := AutenticarUsuario(usuarioValido.Email)

	if usuario.Email != "" {
		t.Fail()
	}

	id := CadastrarUsuario(usuarioValido)

	assertions.NotEqual(id, 0)
}

// Cadastrar usuário
// Usuário inválido - Usuário já está cadastrado no banco de dados
func TestCadastrarUsuario_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{Nome: "Teste Repository", Email: "email1@gmail.com", Senha: "123456"}

	usuario := AutenticarUsuario(usuarioValido.Email)

	if usuario.Email == "" {
		t.Fail()
	}

	id := CadastrarUsuario(usuarioValido)

	assertions.Equal(id, 0)
}

// Obter usuário
// Usuário válido - Retorna o usuário para email e senha informados
func TestObterUsuario_UsuarioValido(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{Email: "email1@gmail.com", Senha: "123456"}

	usuario := ObterUsuario(usuarioValido.Email, usuarioValido.Senha)

	assertions.NotEmpty(usuario)
}

// Obter usuário
// Usuário inválido - Email e senha informados estão incorretos
func TestObterUsuario_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := []model.Usuario{
		{Email: "emai@gmail.com", Senha: "123456"},
		{Email: "email1@gmail.com", Senha: "12345"},
	}

	for _, usuario := range usuarioValido {

		uValido := ObterUsuario(usuario.Email, usuario.Senha)

		assertions.Empty(uValido)
	}
}

// Obter dinheiro inserido
// Usuário válido - Retorna o dinheiro inserido para o usuário informado
func TestObterDinheiroInserido_UsuarioValido(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{ID: 11}

	dinheiroInserido := ObterDinheiroInserido(usuarioValido.ID)

	assertions.NotEmpty(dinheiroInserido)
}

// Obter dinheiro inserido
// Usuário inválido - O id de usuário é inválido para retornar o dinheiro inserido
func TestObterDinheiroInserido_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := model.Usuario{ID: 999}

	dinheiroInserido := ObterDinheiroInserido(usuarioInvalido.ID)

	assertions.Empty(dinheiroInserido)
}

// Obter dinheiro inserido
// Usuário válido mas sem dinheiro - O usuário é válido mas não possui dinheiro registrado
func TestObterDinheiroInserido_UsuarioValidoMasSemDinheiro(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{ID: 12}

	dinheiroInserido := ObterDinheiroInserido(usuarioValido.ID)

	assertions.Empty(dinheiroInserido)
}

// Obter usuário pelo ID
// Usuário válido - Usuário informado está registrado no banco de dados
func TestObterUsuarioPeloID_UsuarioValido(t *testing.T) {

	assertions := require.New(t)

	usuarioValido := model.Usuario{ID: 11}

	dinheiroInserido := ObterUsuarioPeloID(usuarioValido.ID)

	assertions.NotEmpty(dinheiroInserido)
}

// Obter usuário pelo ID
// Usuário inválido - Usuário informado não está registrado no banco de dados
func TestObterUsuarioPeloID_UsuarioInvalido(t *testing.T) {

	assertions := require.New(t)

	usuarioInvalido := model.Usuario{ID: 999}

	dinheiroInserido := ObterUsuarioPeloID(usuarioInvalido.ID)

	assertions.Empty(dinheiroInserido)
}
