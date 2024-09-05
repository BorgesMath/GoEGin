package models

type Meta struct {
	Nome   string  `json:"nome"`
	Check  bool    `json:"Check"`
	Tempo  int64   `json:"Tempo"`
	Passos []Passo `json:"Passo"`
}

var Metas []Meta
