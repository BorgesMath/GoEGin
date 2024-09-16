package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Meta struct {
	gorm.Model
	Nome   string  `json:"nome" validate:"nonzero"`
	Check  bool    `json:"Check" validate:"required"`
	Tempo  int64   `json:"Tempo"`
	Passos []Passo `json:"Passo" gorm:"foreignKey:MetaID;constraint:OnDelete:CASCADE;"`
	// Definir o campo de chave estrangeira MetaID, e fala que quando deletar meta deletar os passos
}

func ValidaDadosMetas(meta *Meta) error {
	if err := validator.Validate(meta); err != nil {
		return err
	}

	return nil
}
