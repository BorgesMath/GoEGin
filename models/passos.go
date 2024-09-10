package models

import "gorm.io/gorm"

type Passo struct {
	gorm.Model
	Nome      string `json:"nomePasso"`
	Check     bool   `json:"CheckPasso"`
	Tempo     int64  `json:"TempoPasso"`
	Descricao string `json:"DescricaoPasso"`
	MetaID    uint   // Campo de chave estrangeira para associar com Meta
}
