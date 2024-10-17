package routes

import (
	ct "github.com/BorgesMath/GoEGin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/metas", ct.ExibeTodasMetas)
	r.GET("/metas/:id", ct.AchaMetaPorID)
	r.GET("/:nome", ct.Inicio)
	r.POST("/metas", ct.CriaNovaMeta)
	r.POST("/passos", ct.CriaNovoPasso)
	r.DELETE("/metas/:id", ct.DeletaMeta)
	r.DELETE("/passos/:idPasso", ct.DeletaPasso)
	r.PATCH("/metas/:id", ct.EditaMeta)
	r.PATCH("/passos/:id", ct.EditaPasso)
	r.GET("/index", ct.ExibePaginaIndex)
	r.NoRoute(ct.RotaNaoEncontrada)
	r.Run(":5000")
	// Subir o servidor,

}
