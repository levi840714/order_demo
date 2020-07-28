package handler

import (
	"order_demo/lib/logger"
	"order_demo/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type GoodStatus struct {
	Status string `form:"status" binding:"required"`
}

func GetGoods(c *gin.Context) {
	var request GoodStatus
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	goods, err := model.GetGoods(request.Status)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Get goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": goods})

}

type Add struct {
	Goods  string  `json:"goods" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func AddGoods(c *gin.Context) {
	var request Add
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	id, err := model.AddGoods(request.Goods, request.Amount)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Add Goods failed", "data": ""})
		return
	}
	c.JSON(201, gin.H{"code": 0, "msg": "", "data": map[string]int{"id": id}})
	return

}

type UpdateId struct {
	ID     int     `json:"id" binding:"required"`
	Goods  string  `json:"goods"`
	Amount float64 `json:"amount"`
}

func UpdateGoods(c *gin.Context) {
	var request UpdateId
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	if err := model.UpdateGoods(request.ID, request.Goods, request.Amount); err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Update goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Update goods success", "data": ""})
}

func DeleteGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	if err := model.DeleteGoods(id); err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Delete goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Delete success", "data": ""})
}

func GetTodaySummary(c *gin.Context) {
	var orderList = make(map[string]map[string]float64)
	orders, err := model.GetTodayOrder(0)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Get today order failed", "data": ""})
		return
	}
	for _, v := range *orders {
		if orderList[v.Account] == nil {
			orderList[v.Account] = make(map[string]float64)
		}
		orderList[v.Account][v.Goods] += v.Amount
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": orderList})
}
