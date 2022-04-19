// Package model contém as entidades usuario e cripto, que representam as entidades do sistema
package model

// Usuario é formado pelo ID, nome, email e senha do usuário
type Usuario struct {
	ID    int    `json:"user_id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
