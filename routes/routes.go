package routes

import (
	ct "github.com/BorgesMath/GoEGin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/metas", ct.ExibeTodasMetas)
	r.GET("/:nome", ct.Inicio)
	r.POST("/metas", ct.CriaNovaMeta)
	r.POST("/passos", ct.CriaNovoPasso)
	r.Run(":5000")
	// Subir o servidor,

}
