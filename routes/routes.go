package routes

import (
	ct "github.com/BorgesMath/GoEGin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/metas", ct.ExibeTodasMetas)
	r.GET("/metas/:id", ct.AchaMetaPorID)
	r.GET("/:nome", ct.Inicio)
	r.POST("/metas", ct.CriaNovaMeta)
	r.POST("/passos", ct.CriaNovoPasso)
	r.DELETE("/metas/:id", ct.DeletaMeta)
	r.DELETE("/metas/:idMeta/passos/:idPasso", ct.DeletaPasso)
	r.PATCH("/metas/:id", ct.EditaMeta)
	r.PATCH("/passos/:id", ct.EditaPasso)
	r.Run(":5000")
	// Subir o servidor,

}
