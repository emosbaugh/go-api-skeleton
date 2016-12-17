package errors

import "gopkg.in/gin-gonic/gin.v1"

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors.ByType(gin.ErrorTypePublic)
		if len(errs) > 0 {
			c.JSON(-1, gin.H{"errors": errs})
		}
	}
}
