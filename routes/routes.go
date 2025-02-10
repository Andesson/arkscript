package router

import (
	"github.com/Andesson/arkscript/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/comands")
	{
		v1.GET("/dino/import", handler.ImportDinos)
		//v1.POST("dino/save", handler.CreateOpeningHandler)
		//v1.DELETE("dino/delete", handler.DeleteOpeningHandler)
		//v1.PUT("dino/update", handler.UpdateOpeningHandler)
		//v1.GET("/dinos", handler.ShowAllOpeningHandler)
	}
}
