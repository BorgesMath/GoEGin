package models

import "gorm.io/gorm"

type Meta struct {
	gorm.Model
	Nome   string  `json:"nome"`
	Check  bool    `json:"Check"`
	Tempo  int64   `json:"Tempo"`
	Passos []Passo `json:"Passo" gorm:"foreignKey:MetaID;constraint:OnDelete:CASCADE;"`
	// Definir o campo de chave estrangeira MetaID
}
