package handler

import (
	"order_demo/lib/logger"
	"order_demo/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type DepositData struct {
	AccountId int     `json:"column:id"`
	Amount    float64 `json:"amount" binding:"required"`
}

func Deposit(c *gin.Context) {
	var request DepositData
	accountId := c.MustGet("accountId").(int)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		logger.Error.Println(err.Error())
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}
	request.AccountId = accountId
	result := model.AddTransfer(request.AccountId, request.Amount)
	if !result {
		c.JSON(500, gin.H{"code": 1, "msg": "Transfer failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "Transfer success", "data": ""})
}
