package repository

import (
	"minhascriptos/model"
	"minhascriptos/repository"
	"testing"
)

type autenticarUsuario struct {
	email   string
	usuario model.Usuario
}

var autenticarUsuarios = []autenticarUsuario{
	{"alberto@gmail.com", model.Usuario{ID: 1, Nome: "Alberto", Email: "alberto@gmail.com", Senha: "123456"}},
	{"test@gmail.com", model.Usuario{}},
}

func TestAutenticarUsuario(t *testing.T) {
	for _, teste := range autenticarUsuarios {
		resultado := repository.AutenticarUsuario(teste.email)
		if resultado != teste.usuario {
			t.Fatal("Informação dada:", teste.email, "| Informação esperada:", teste.usuario,
				"| Informação recebida:", resultado)
		}
	}
}

type cadastrarUsuario struct {
	usuario model.Usuario
	id      int
}

var cadastrarUsuarios = []cadastrarUsuario{
	{model.Usuario{ID: 0, Nome: "Mariana", Email: "mariana@gmail.com", Senha: "123456"}, 0},
	{model.Usuario{ID: 1, Nome: "Alberto", Email: "alberto@gmail.com", Senha: "123456"}, 0},
}

func TestCadastrarUsuario(t *testing.T) {
	for _, teste := range cadastrarUsuarios {
		resultado := repository.CadastrarUsuario(teste.usuario)
		if teste.usuario.Email != "alberto@gmail.com" {
			if resultado <= 0 {
				t.Fatal("Informação dada:", teste.usuario, "| Informação esperada:", teste.id,
					"| Informação recebida:", resultado)
			}
		} else {
			if resultado != 0 {
				t.Fatal("Informação dada:", teste.usuario, "| Informação esperada:", teste.id,
					"| Informação recebida:", resultado)
			}
		}
	}
}
