package controllers

import (
	"net/http"

	"github.com/BorgesMath/GoEGin/database"
	"github.com/BorgesMath/GoEGin/models"
	"github.com/gin-gonic/gin"
)

func ExibePaginaIndex(c *gin.Context) {

	var metas []models.Meta
	database.DB.Preload("Passos").Find(&metas)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"metas": metas,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
