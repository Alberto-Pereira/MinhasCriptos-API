package model

type Cripto struct {
	ID                          int     `json:"cripto_id"`
	TipoMoeda                   string  `json:"tipo_moeda"`
	DataDeCompra                string  `json:"data_de_compra"`
	QuantidadeComprada          float64 `json:"quantidade_de_compra"`
	PrecoDeCompra               float64 `json:"preco_de_compra"`
	ValorDaUnidadeNoDiaDeCompra float64 `json:"valor_da_unidade_no_dia_de_compra"`
	UsuarioId                   Usuario `json:"usuario_id"`
}
