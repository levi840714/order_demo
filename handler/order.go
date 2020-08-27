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

// @Summary 充值
// @Tags User
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Param payload body DepositData true "充值金額"
// @Success 200 {string} json "{"code": 0, "msg": "Transfer success!", "data": {"transferId": {充值單號}, "balance": {目前餘額}}"
// @Router /api/deposit [post]
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

// @Summary 送出訂單
// @Tags Order
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Param payload body OrderData true "餐點ID"
// @Success 200 {string} json "{"code": 0, "msg": "", "data": "{"id": 訂單號}"}"
// @Router /api/order [post]
func NewOrder(c *gin.Context) {
	var request OrderData
	accountId := c.MustGet("accountId").(int)
	if err := c.ShouldBindWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}

	accountInfo, e := model.GetAccountInfo(accountId)
	GoodsInfo, err := model.GetGoodsInfo(request.GoodsId)
	if e != nil || err != nil {
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Order failed", "data": ""})
		return
	}

	if accountInfo.Balance < GoodsInfo.Amount {
		c.JSON(400, gin.H{"code": 1, "msg": "Balance not enough", "data": ""})
		return
	}

	model.TX = model.DB.Begin()
	id, err := model.NewOrder(accountId, request.GoodsId)
	if err != nil {
		model.TX.Rollback()
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Order failed", "data": ""})
		return
	}

	if _, err := model.UpdateBalance(accountId, GoodsInfo.Amount*-1); err != nil {
		model.TX.Rollback()
		logger.Error.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Order failed", "data": ""})
		return
	}

	model.TX.Commit()
	c.JSON(201, gin.H{"code": 0, "msg": "", "data": map[string]int{"id": id}})
}

// @Summary 刪除訂單
// @Tags Order
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Param id path string true "訂單ID"
// @Success 200 {string} json "{"code": 0, "msg": "Delete success", "data": ""}"
// @Router /api/order/{id} [delete]
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

// @Summary 取得今日訂單
// @Tags Order
// @version 1.0
// @Accpet json
// @Produce json
// @Security ApiToken
// @Success 200 {string} json "{"code": 0, "msg": "", "data": "{訂單}"}"
// @Router /api/order [get]
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
