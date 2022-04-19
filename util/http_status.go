// Package util contém arquivos que auxiliam no desenvolvimento do sistema
package util

// HttpStatus é formado por um ID e Mensagem que representam
// o número de uma resposta http e sua mensagem associada
type HttpStatus struct {
	ID       int
	Mensagem string
}
