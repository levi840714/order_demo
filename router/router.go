package router

import (
	"order_demo/handler"
	"order_demo/lib/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
	api := router.Group("/api")
	api.Use(middleware.CheckJwtValid)
	{
		api.POST("/deposit", handler.Deposit)
	}
	return router

}
