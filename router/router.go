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
	router.GET("/goods", handler.GetGoods)

	api := router.Group("/api")
	admin := router.Group("/admin")

	api.Use(middleware.CheckJwtValid)
	{
		api.POST("/deposit", handler.Deposit)
		api.GET("/order", handler.GetTodayOrder)
		api.POST("/order", handler.NewOrder)
		api.DELETE("/order/:id", handler.DeleteOrder)
	}

	admin.Use(middleware.CheckJwtValid, middleware.CheckAdmin)
	{
		admin.POST("/goods", handler.AddGoods)
		admin.PUT("/goods", handler.UpdateGoods)
		admin.DELETE("/goods/:id", handler.DeleteGoods)
		admin.GET("/summaryList", handler.GetTodaySummary)
	}

	return router
}
