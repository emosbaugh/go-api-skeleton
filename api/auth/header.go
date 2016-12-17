package auth

import "gopkg.in/gin-gonic/gin.v1"

func Header(c *gin.Context, token string) {
	c.Header("Authorization", "Bearer "+token)
}
