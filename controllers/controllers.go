package controllers

import (
	"log"
	"net/http"

	"github.com/BorgesMath/GoEGin/database"
	"github.com/BorgesMath/GoEGin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodasMetas(c *gin.Context) {
	// para o get a função precisa apontar para um
	//gin.Context

	var metas []models.Meta
	database.DB.Preload("Passos").Find(&metas) // nome tem que ser similar ao db
	//Preload para carregar a conexão etre eles antes
	// 200 eh o httpCode
	c.JSON(200, metas)

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

	if err := models.ValidaDadosMetas(&meta); err != nil {
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

	// APAGAR ESSA PARTE E FAZER PELO VALIDATOR!!!
	// Verifica se o MetaID foi passado
	if passo.MetaID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "MetaID é obrigatório para criar um Passo",
		})
		return
	}

	if err := models.ValidaDadosPassos(&passo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&passo)
	c.JSON(http.StatusCreated, passo) // Retorna 201
}

func AchaMetaPorID(c *gin.Context) {
	var meta models.Meta
	id := c.Params.ByName("id")

	database.DB.Preload("Passos").First(&meta, id) // confuso como ele sabe relacionar
	// mas é pelo nome da variavel

	if meta.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Meta não existe"})

		return
	}

	c.JSON(http.StatusOK, meta)

}

func DeletaMeta(c *gin.Context) {
	var meta models.Meta
	id := c.Param("id")
	database.DB.Delete(&meta, id)
	c.JSON(http.StatusOK, gin.H{"message": "Meta deletado com sucesso!"})

}

func DeletaPasso(c *gin.Context) {
	// Acho que minha logica foi errada, não preciso do idMeta
	// deixar pra se precisar futuramente
	var passo models.Passo
	idMeta := c.Param("idMeta")
	idPasso := c.Param("idPasso")

	// Log para verificação
	log.Printf("Meta ID: %s, Passo ID: %s", idMeta, idPasso)

	database.DB.Delete(&passo, idPasso)

	c.JSON(200, gin.H{
		"message": "Passo deletado com sucesso!",
	})
}

func EditaMeta(c *gin.Context) {
	var meta models.Meta
	id := c.Params.ByName("id")

	database.DB.First(&meta, id)

	if err := c.ShouldBindBodyWithJSON(&meta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&meta).UpdateColumns(meta)
	c.JSON(http.StatusOK, meta)
}

func EditaPasso(c *gin.Context) {

	var passo models.Passo
	id := c.Params.ByName("id")

	database.DB.First(&passo, id)

	if err := c.ShouldBindBodyWithJSON(&passo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&passo).UpdateColumns(passo)
	c.JSON(http.StatusOK, passo)

}
