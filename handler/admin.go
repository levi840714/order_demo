package handler

import (
	"order_demo/lib/logger"
	"order_demo/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Goods struct {
}

func GetGoods(c *gin.Context) {

}

type Add struct {
	Good   string  `json:"good" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func AddGoods(c *gin.Context) {
	var request Add
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	id, err := model.AddGoods(request.Good, request.Amount)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Add Goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": map[string]int{"id": id}})
	return

}

func UpdateGoods(c *gin.Context) {

}

type Delete struct {
	ID     int    `json:"id"  binding:"required"`
	Status string `json:"status"  binding:"required"`
}

func DeleteGoods(c *gin.Context) {
	var request Delete
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}
	if err := model.DeleteGoods(request.ID, request.Status); err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "Delete goods failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Delete success", "data": ""})
}
