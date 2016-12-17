package auth

import (
	"fmt"
	"net/http"

	"github.com/replicatedcom/gin-example/api/errors"
	"github.com/replicatedcom/gin-example/logging"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"gopkg.in/gin-gonic/gin.v1"
)

var ErrNoToken = fmt.Errorf("token does not exist in context")

func Middleware(secret string, logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var claims jwt.StandardClaims
		_, err := request.ParseFromRequestWithClaims(c.Request, request.AuthorizationHeaderExtractor, &claims, keyFunc(secret))
		if err != nil {
			if logger != nil {
				logger.Debugf("failed to parse jwt token from authorization request header: %v", err)
			}
			if err == request.ErrNoTokenInRequest {
				errors.New(http.StatusUnauthorized, "header.authorization.empty", "token not found in authorization header", nil).Response(c)
				return
			}
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("claims", claims)
	}
}

func GetClaimsFromGinContext(c *gin.Context) (*jwt.StandardClaims, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, ErrNoToken
	}
	standardClaims, ok := claims.(jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected claims type %T", claims)
	}
	return &standardClaims, nil
}
