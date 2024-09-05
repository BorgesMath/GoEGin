package controllers

import (
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
