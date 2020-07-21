package router

import (
	"order_demo/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", handler.Login)
		api.POST("/register", handler.Register)
	}
	return router

}
