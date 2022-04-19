// Package service contém as regras de serviço das entidades usuário e cripto
package service

import (
	"minhascriptos/model"
	"minhascriptos/repository"
	"minhascriptos/util"
	"regexp"
)

// Cadastrar Usuario
// A entidade usuário, após ser validada, é enviada para o repositório para ser cadastrada
// Se ela for cadastrada, é retornado true e o status associado
// Se ela não for cadastrada, é retornado false e o status associado
func CadastrarUsuario(usuario model.Usuario) (bool, util.HttpStatus) {

	if isUsuarioValido(usuario) == false {
		return false, util.HttpStatus{ID: 400, Mensagem: "Usuário inválido!"}
	}

	if isUsuarioCadastrado(usuario.Email) == true {
		return false, util.HttpStatus{ID: 406, Mensagem: "Usuário já cadastrado!"}
	}

	usuarioCadastradoID := repository.CadastrarUsuario(usuario)

	if usuarioCadastradoID != 0 {
		return true, util.HttpStatus{ID: 200, Mensagem: "Usuário cadastrado!"}
	} else {
		return false, util.HttpStatus{ID: 500, Mensagem: "Erro desconhecido ao tentar cadastrar usuário!"}
	}

}

// Obter Usuario
// É fornecido email e senha, que após serem validados, são enviados para o repositório para buscar um usuário
// Se ele for buscado, é retornado o usuário correspondente, true e o status associado
// Se ele não for buscado, é retornado nil, false e o status associado
func ObterUsuario(email string, senha string) (model.Usuario, bool, util.HttpStatus) {

	if isEmailESenhaValida(email, senha) == false {
		return model.Usuario{}, false, util.HttpStatus{ID: 400,
			Mensagem: "Email ou senha não seguem os padrões requisitados!"}
	}

	if isUsuarioCadastrado(email) == true {

		usuario := repository.ObterUsuario(email, senha)

		if usuario.ID == 0 {
			return model.Usuario{}, false, util.HttpStatus{ID: 406, Mensagem: "Senha inválida para o email informado!"}

		} else {

			if isUsuarioValido(usuario) == true {
				return usuario, true, util.HttpStatus{ID: 200}
			} else {
				return model.Usuario{}, false, util.HttpStatus{ID: 500,
					Mensagem: "Houve um problema ao validar usuário requisitado!"}
			}
		}

	} else {
		return model.Usuario{}, false, util.HttpStatus{ID: 404,
			Mensagem: "Usuário não encontrado para o email informado!"}
	}

}

// Obter Dinheiro Inserido
// A entidade usuário, após ser validada, é enviada para o repositório para obter o dinheiro inserido
// Se a entidade existir e ter dinheiro inserido, é retornado o dinheiro inserido correspondente, true e o status associado
// Se a entidade existir e não ter dinheiro inserido, é retornado nil, false e o status associado
func ObterDinheiroInserido(usuario_id int) ([]model.DinheiroInserido, bool, util.HttpStatus) {

	if isUsuarioIDValido(usuario_id) == false {
		return []model.DinheiroInserido{}, false, util.HttpStatus{ID: 400, Mensagem: "ID inválido!"}
	}

	dinheiroInserido := repository.ObterDinheiroInserido(usuario_id)

	if dinheiroInserido == nil {
		return dinheiroInserido, false, util.HttpStatus{ID: 404,
			Mensagem: "Dinheiro inserido não encontrado para o ID informado!"}
	} else {
		return dinheiroInserido, true, util.HttpStatus{ID: 200}
	}
}

// Is Usuario Cadastrado
// O email fornecido, é enviado para o repositório
// Se ele for autenticado, é retornado true
// Se ele não for autenticado, é retornado false
func isUsuarioCadastrado(email string) bool {

	usuario := repository.AutenticarUsuario(email)

	if usuario.Email != email {
		return false
	} else {
		return true
	}
}

// Is Usuario Valido
// A entidade usuário fornecida, é validada pelas regras definidas
// Se ela for validada, é retornado true
// Se ela não for validada, é retornado false
func isUsuarioValido(usuario model.Usuario) bool {

	// ! Formato aceito - Primeira letra maiúscula, pelo menos duas letras e aceita espaço
	nomeValido := regexp.MustCompile(`^([A-Z]{1}[a-z]+\s?)+$`)
	// ! Formato aceito - exemplo@email.com
	emailValido := regexp.MustCompile(`^[a-z0-9._]+[@]{1}[a-z0-9]+[.]{1}[a-z]+$`)
	// ! Formato aceito - Deve conter pelo menos 3 caracteres
	senhaValida := regexp.MustCompile(`^[.\s\S\d\D\w\W]{3,}$`)

	if nomeValido.MatchString(usuario.Nome) {
		if emailValido.MatchString(usuario.Email) {
			if senhaValida.MatchString(usuario.Senha) {
				return true
			}
		}
	}

	return false
}

// Is Email E Senha Valida
// O email e senha fornecidos, são validados pelas regras definidas
// Se eles forem validados, é retornado true
// Se eles não forem validados, é retornado false
func isEmailESenhaValida(email string, senha string) bool {

	// ! Formato aceito - exemplo@email.com
	emailValido := regexp.MustCompile(`^[a-z0-9._]+[@]{1}[a-z0-9]+[.]{1}[a-z]+$`)
	// ! Formato aceito - Deve conter pelo menos 3 caracteres
	senhaValida := regexp.MustCompile(`^[.\s\S\d\D\w\W]{3,}$`)

	if emailValido.MatchString(email) {
		if senhaValida.MatchString(senha) {
			return true
		}
	}

	return false
}

// Is Usuario ID Valido
// É fornecido um id de usuário para ser buscado
// Se ele for encontrado, é retornado true
// Se ele não for encontrado, é retornado false
func isUsuarioIDValido(usuario_id int) bool {

	u := repository.ObterUsuarioPeloID(usuario_id)

	if u != usuario_id {
		return false
	} else {
		return true
	}
}
