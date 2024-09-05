package main

import (
	"github.com/BorgesMath/GoEGin/models"
	"github.com/BorgesMath/GoEGin/routes"
)

func main() {

	models.Metas = []models.Meta{
		{
			Nome:  "Meta de Aprender Golang",
			Check: false,
			Tempo: 12000,
			Passos: []models.Passo{
				{
					Nome:      "Instalar Golang",
					Check:     true,
					Tempo:     300,
					Descricao: "Baixar e instalar Golang no sistema.",
				},
				{
					Nome:      "Configurar Ambiente",
					Check:     false,
					Tempo:     600,
					Descricao: "Configurar o ambiente de desenvolvimento com um editor de texto e Go.",
				},
			},
		},
		{
			Nome:  "Meta de Ler Livros",
			Check: false,
			Tempo: 25000,
			Passos: []models.Passo{
				{
					Nome:      "Escolher Livros",
					Check:     true,
					Tempo:     200,
					Descricao: "Selecionar uma lista de livros para ler.",
				},
				{
					Nome:      "Ler 1º Livro",
					Check:     false,
					Tempo:     7200,
					Descricao: "Ler o primeiro livro da lista.",
				},
				{
					Nome:      "Ler 2º Livro",
					Check:     false,
					Tempo:     7200,
					Descricao: "Ler o segundo livro da lista.",
				},
			},
		},
		{
			Nome:  "Meta de Exercícios Físicos",
			Check: false,
			Tempo: 18000,
			Passos: []models.Passo{
				{
					Nome:      "Fazer Aquecimento",
					Check:     true,
					Tempo:     600,
					Descricao: "Realizar 10 minutos de aquecimento.",
				},
				{
					Nome:      "Exercício Cardio",
					Check:     false,
					Tempo:     3600,
					Descricao: "Fazer 1 hora de corrida ou bicicleta.",
				},
			},
		},
		{
			Nome:  "Meta de Aprender Cozinhar",
			Check: false,
			Tempo: 10000,
			Passos: []models.Passo{
				{
					Nome:      "Assistir Aulas",
					Check:     true,
					Tempo:     2000,
					Descricao: "Assistir às aulas de culinária.",
				},
				{
					Nome:      "Fazer Pratos Simples",
					Check:     false,
					Tempo:     5000,
					Descricao: "Cozinhar pratos básicos como arroz e feijão.",
				},
			},
		},
		{
			Nome:  "Meta de Meditação",
			Check: false,
			Tempo: 15000,
			Passos: []models.Passo{
				{
					Nome:      "Escolher Local",
					Check:     true,
					Tempo:     300,
					Descricao: "Escolher um local tranquilo para meditar.",
				},
				{
					Nome:      "Meditar Diariamente",
					Check:     false,
					Tempo:     14400,
					Descricao: "Meditar por 30 minutos todos os dias durante um mês.",
				},
			},
		},
	}

	routes.HandleRequests()

}
