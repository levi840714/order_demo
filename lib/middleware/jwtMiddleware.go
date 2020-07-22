package middleware

import (
	"order_demo/lib/auth"

	"github.com/gin-gonic/gin"
)

func CheckJwtValid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(400, gin.H{"code": 1, "msg": "Please login! ", "data": ""})
		c.Abort()
		return
	}
	claims, err := auth.Verify(authToken)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error(), "data": ""})
		c.Abort()
		return
	}
	c.Set("role", claims.Role)
	c.Set("accountId", claims.AccountId)
	c.Next()
}
