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
	admin := router.Group("/admin")
	api := router.Group("/api")

	api.Use(middleware.CheckJwtValid)
	{
		api.POST("/deposit", handler.Deposit)
	}

	admin.Use(middleware.CheckJwtValid, middleware.CheckAdmin)
	{
		admin.GET("/goods", handler.GetGoods)
		admin.POST("/goods", handler.AddGoods)
		admin.PUT("/goods/:id", handler.UpdateGoods)
		admin.DELETE("/goods/:id", handler.DelGoods)
	}

	return router
}
