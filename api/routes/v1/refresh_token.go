package v1

import (
	"net/http"

	"github.com/replicatedcom/gin-example/api/auth"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

func RefreshToken(c *gin.Context) {
	claims, err := auth.GetClaimsFromGinContext(c)
	if err != nil {
		log.WithField("err", err).Error("failed to get token claims from context")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	token, err := auth.NewToken(claims.Id, viper.GetString("secret"))
	if err != nil {
		log.WithField("err", err).Error("failed to generate token")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	auth.Header(c, token)
	c.Status(http.StatusNoContent)
}
