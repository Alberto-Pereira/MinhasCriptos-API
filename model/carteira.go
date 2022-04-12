package model

type DinheiroInserido struct {
	TipoMoeda string  `json:"user_id"`
	Total     float64 `json:"total"`
}

type DinheiroInseridoSlice []DinheiroInserido
