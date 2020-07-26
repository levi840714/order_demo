package handler

import (
	"order_demo/lib/logger"
	"order_demo/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type DepositData struct {
	Amount float64 `json:"amount" binding:"required"`
}

func Deposit(c *gin.Context) {
	var request DepositData
	accountId := c.MustGet("accountId").(int)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
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
