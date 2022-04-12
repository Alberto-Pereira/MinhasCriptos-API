package service

import (
	"fmt"
	"minhascriptos/model"
	"minhascriptos/repository"
)

func autenticarUsuario(email string) bool {
	usuario := repository.AutenticarUsuario(email)

	fmt.Println("EMAIL INFORMADO =>", email)

	if usuario.Email != email {
		fmt.Println("Usuário não encontrado para o EMAIL informado!")
		return false
	} else {
		fmt.Println("Usuário encontrado para o EMAIL informado!")
		return true
	}

}

func CadastrarUsuario(usuario model.Usuario, email string) bool {
	usuarioAutenticado := autenticarUsuario(email)

	if usuarioAutenticado == true {
		fmt.Println("Usuário já cadastrado!")
		return false
	} else {
		fmt.Println("Cadastrando usuário...")

		id := repository.CadastrarUsuario(usuario)

		if id != 0 {
			fmt.Println("Usuário cadastrado!")
			return true
		} else {
			fmt.Println("Erro ao cadastrar usuário!")
			return false
		}
	}
}

func ObterUsuario(email string, senha string) (model.Usuario, bool) {
	isEmailValido := autenticarUsuario(email)

	if isEmailValido == true {
		usuario := repository.ObterUsuario(email, senha)

		if usuario.ID == 0 {
			fmt.Println("Senha inválida!")
		}

		return usuario, true

	} else {
		fmt.Println("Email inválido! =>", email)
	}

	return model.Usuario{}, false
}

func ObterDinheiroInserido(usuario_id int) model.DinheiroInseridoSlice {
	dis := model.DinheiroInseridoSlice{}

	dis = repository.ObterDinheiroInserido(usuario_id)

	if dis == nil {
		fmt.Println("Obter Dinheiro Inserido | Slice vazio ao obter dinheiro inserido com ID:", usuario_id)
		return dis
	} else {
		fmt.Println("Obter Dinheiro Inserido | Slice encontrado para o ID:", usuario_id)
		return dis
	}
}
