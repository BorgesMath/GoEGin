package routes

import (
	ct "github.com/BorgesMath/GoEGin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/metas", ct.ExibeTodasMetas)
	r.GET("/:nome", ct.Inicio)
	r.Run(":5000")
	// Subir o servidor,

}
