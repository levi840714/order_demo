package handler

import (
	"order_demo/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RegisterAccount struct {
	Account  string `json:"account" binding:"required"`  // 帳號
	Password string `json:"password" binding:"required"` // 密碼
}

func Register(c *gin.Context) {
	var request RegisterAccount
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}
	code, err := model.RegisterAccount(request.Account, request.Password)
	if code != 0 {
		c.JSON(500, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "register success", "data": ""})
}

type LoginAccount struct {
	Account  string `json:"account" binding:"required"`  // 帳號
	Password string `json:"password" binding:"required"` // 密碼
}

func Login(c *gin.Context) {
	var request LoginAccount
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		return
	}
	check := model.CheckLogin(request.Account, request.Password)
	if !check {
		c.JSON(500, gin.H{"code": 1, "msg": "login failed", "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "login success", "data": ""})
}
