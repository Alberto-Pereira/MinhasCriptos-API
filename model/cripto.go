// Package model contém as entidades usuario e cripto, que representam as entidades do sistema
package model

// Cripto é formado pelo ID
// TipoMoeda representa o tipo de moeda
// DataDeCompra representa a data de compra da moeda
// QuantidadeComprada representa a quantidade daquele tipo de moeda
// PrecoDeCompra representa o preco de compra daquele tipo de moeda
// ValorDaUnidadeNoDiaDeCompra representa o valor de uma unidade no dia da compra
// UsuarioID representa o id do usuario associado a compra da moeda
type Cripto struct {
	ID                          int     `json:"cripto_id"`
	TipoMoeda                   string  `json:"tipo_moeda"`
	DataDeCompra                string  `json:"data_de_compra"`
	QuantidadeComprada          float64 `json:"quantidade_comprada"`
	PrecoDeCompra               float64 `json:"preco_de_compra"`
	ValorDaUnidadeNoDiaDeCompra float64 `json:"valor_da_unidade_no_dia_de_compra"`
	UsuarioId                   Usuario `json:"usuario_id"`
}

// DinheiroInserido é uma estrutura criada para auxiliar no retorno das moedas do usuário
// TipoMoeda representa o tipo de moeda
// Total representa o total para aquele tipo de moeda
type DinheiroInserido struct {
	TipoMoeda string  `json:"tipo_moeda"`
	Total     float64 `json:"total"`
}
