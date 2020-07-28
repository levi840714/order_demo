package handler

import (
	"order_demo/lib/logger"
	"order_demo/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type DepositData struct {
	Amount float64 `json:"amount" binding:"required"`
}

func Deposit(c *gin.Context) {
	var request DepositData
	accountId := c.MustGet("accountId").(int)
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	model.TX = model.DB.Begin()
	transferId, err := model.AddTransfer(accountId, request.Amount)
	if err != nil {
		model.TX.Rollback()
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Transfer failed", "data": ""})
		return
	}

	balance, err := model.UpdateBalance(accountId, request.Amount)
	if err != nil {
		model.TX.Rollback()
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Transfer failed", "data": ""})
		return
	}
	model.TX.Commit()

	successData := map[string]interface{}{
		"transferId": transferId,
		"balance":    balance,
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Transfer success!", "data": successData})
}

type OrderData struct {
	GoodsId int `json:"goodsId" binding:"required"`
}

func NewOrder(c *gin.Context) {
	var request OrderData
	accountId := c.MustGet("accountId").(int)
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	id, err := model.NewOrder(accountId, request.GoodsId)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Order failed", "data": ""})
		return
	}
	c.JSON(201, gin.H{"code": 0, "msg": "", "data": map[string]int{"id": id}})
}

func DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	if err := model.DeleteOrder(id); err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Delete order failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Delete success", "data": ""})
}

func GetTodayOrder(c *gin.Context) {
	accountId := c.MustGet("accountId").(int)
	orders, err := model.GetTodayOrder(accountId)
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Get today order failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": orders})
}
