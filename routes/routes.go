package router

import (
	"github.com/Andesson/arkscript/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/comands")
	{
		v1.GET("dino/import", handler.ImportDinos)
		v1.GET("item/import", handler.ImportItens)
		v1.GET("chibi/import", handler.ImportChibis)
		v1.GET("all/import", handler.ImportAll)

		v1.POST("dino/spawn", handler.GeraScriptDino)
		v1.POST("item/spawn", handler.GeraScriptItem)
		v1.POST("chibi/spawn", handler.GeraScriptChibi)
	}
}
