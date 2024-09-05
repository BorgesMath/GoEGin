package models

type Passo struct {
	Nome      string `json:"nomePasso"`
	Check     bool   `json:"CheckPasso"`
	Tempo     int64  `json:"TempoPasso"`
	Descricao string `json:"DescricaoPasso"`
}
