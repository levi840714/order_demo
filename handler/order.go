package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Deposit(c *gin.Context) {
	fmt.Println(c.MustGet("accountId"), c.MustGet("role"))
}
