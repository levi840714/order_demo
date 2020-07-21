package router

import (
	"order_demo/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()
	router.POST("/login", handler.Login)
	api := router.Group("/api")
	{
		api.POST("/register", handler.Register)
	}
	return router

}
