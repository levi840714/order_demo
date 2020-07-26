package middleware

import "github.com/gin-gonic/gin"

func CheckAdmin(c *gin.Context) {
	role := c.MustGet("role")
	if role != 1 {
		c.Abort()
		c.JSON(401, gin.H{"code": 1, "msg": "permission denied", "data": ""})
		return
	}
	c.Next()
}
