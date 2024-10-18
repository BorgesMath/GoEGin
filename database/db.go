package database

import (
	"log"

	"github.com/BorgesMath/GoEGin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: " + err.Error())
	}
	DB.AutoMigrate(&models.Meta{}, &models.Passo{}) // Migrar ambos os modelos
}
