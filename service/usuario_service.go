package service

import (
	"minhascriptos/model"
	"minhascriptos/repository"
	"minhascriptos/util"
	"regexp"
)

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

func isUsuarioCadastrado(email string) bool {

	usuario := repository.AutenticarUsuario(email)

	if usuario.Email != email {
		return false
	} else {
		return true
	}
}

func isUsuarioValido(usuario model.Usuario) bool {

	nomeValido := regexp.MustCompile(`^([A-Z]{1}[a-z]+\s?)+$`)
	emailValido := regexp.MustCompile(`^[a-z0-9._]+[@]{1}[a-z0-9]+[.]{1}[a-z]+$`)
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

func isEmailESenhaValida(email string, senha string) bool {

	emailValido := regexp.MustCompile(`^[a-z0-9._]+[@]{1}[a-z0-9]+[.]{1}[a-z]+$`)
	senhaValida := regexp.MustCompile(`^[.\s\S\d\D\w\W]{3,}$`)

	if emailValido.MatchString(email) {
		if senhaValida.MatchString(senha) {
			return true
		}
	}

	return false
}

func isUsuarioIDValido(usuario_id int) bool {

	u := repository.ObterUsuarioPeloID(usuario_id)

	if u != usuario_id {
		return false
	} else {
		return true
	}
}
