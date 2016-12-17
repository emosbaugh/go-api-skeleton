package context

import (
	"github.com/replicatedcom/gin-example/api/auth"
	"github.com/replicatedcom/gin-example/logging"

	"gopkg.in/gin-gonic/gin.v1"
)

func Middleware(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := auth.GetClaimsFromGinContext(c)
		if err != nil {
			if err != auth.ErrNoToken && logger != nil {
				logger.Warningf("failed to get jwt claims from request header: %v", err)
			}
			return
		}
		// TODO: is this a User or APIToken?
		SetUserID(c, claims.Id)
	}
}
