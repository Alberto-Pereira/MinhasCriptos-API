package model

type Usuario struct {
	ID    int    `json:"user_id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
