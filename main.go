package main

import (
	"github.com/gin-gonic/gin"
)

func ExibeTodasMetas(c *gin.Context) {
	// para o get a função precisa apontar para um
	//gin.Context

	// 200 eh o httpCode
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "ED5",
	})

}

func main() {

	r := gin.Default()
	r.GET("/metas", ExibeTodasMetas)
	r.Run(":5000")
	// Subir o servidor,

}
