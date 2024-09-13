package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Passo struct {
	gorm.Model
	Nome      string `json:"nomePasso" validate:"nonzero"`
	Check     bool   `json:"CheckPasso"`
	Tempo     int64  `json:"TempoPasso"`
	Descricao string `json:"DescricaoPasso"`
	MetaID    uint   `json:"MetaID" validate:"nonzero"` // Campo de chave estrangeira para associar com Meta
}

func ValidaDadosPassos(passo *Passo) error {
	if err := validator.Validate(passo); err != nil {
		return err
	}

	return nil
}
