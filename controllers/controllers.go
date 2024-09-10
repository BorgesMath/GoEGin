package controllers

import (
	"net/http"

	"github.com/BorgesMath/GoEGin/database"
	"github.com/BorgesMath/GoEGin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodasMetas(c *gin.Context) {
	// para o get a função precisa apontar para um
	//gin.Context

	// 200 eh o httpCode
	c.JSON(200, models.Metas)

}

func Inicio(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API": "Eai " + nome + " Bom dia",
	})
}

func CriaNovaMeta(c *gin.Context) {
	var meta models.Meta
	if err := c.ShouldBindJSON(&meta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&meta)
	c.JSON(http.StatusCreated, meta) // Melhor que StatusOK?

}

func CriaNovoPasso(c *gin.Context) {
	var passo models.Passo
	if err := c.ShouldBindJSON(&passo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Verifica se o MetaID foi passado
	if passo.MetaID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "MetaID é obrigatório para criar um Passo",
		})
		return
	}

	database.DB.Create(&passo)
	c.JSON(http.StatusCreated, passo) // Retorna 201
}
