package v1

import (
	"database/sql"
	"net/http"

	"github.com/replicatedcom/gin-example/api/auth"
	"github.com/replicatedcom/gin-example/api/context"
	"github.com/replicatedcom/gin-example/api/errors"
	"github.com/replicatedcom/gin-example/api/routes/v1/params"
	"github.com/replicatedcom/gin-example/api/routes/v1/resource"
	"github.com/replicatedcom/gin-example/api/routes/v1/schema"
	"github.com/replicatedcom/gin-example/api/schemas"
	"github.com/replicatedcom/gin-example/inject"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/gin-gonic/gin.v1"
)

var userLoginParamsSchema *gojsonschema.Schema

func init() {
	userLoginParamsSchema = schema.MustLoadSchema("file:///user_login_params.json")
}

func UserLogin(env *inject.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		schemas.LinkHeader(c, "schema.json#/properties/loginParams")

		var request params.UserLogin
		err := schemas.Bind(c, &request, userLoginParamsSchema)
		if err != nil {
			log.WithField("err", err).Error("failed to bind to login params schema")
			return
		}

		user, err := env.UserService.Get(request.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				errors.ErrLoginFailure.Response(c) // hide actual error for security purposes
				return
			}
			log.WithField("err", err).Error("failed to retrieve user")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		err = user.VerifyPassword(request.Password)
		if err != nil {
			log.WithField("err", err).Error("failed to verify password")
			errors.ErrLoginFailure.Response(c) // hide actual error for security purposes
			return
		}

		context.SetUserID(c, user.Email)

		token, err := auth.NewToken(user.Email, viper.GetString("secret"))
		if err != nil {
			log.WithField("err", err).Error("failed to generate token")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		resource := resource.NewUser(user)
		err = schemas.MaybeValidateResponse(c, resource, userResourceSchema)
		if err != nil {
			log.WithField("err", err).Error("failed to validate resource")
			return
		}
		auth.Header(c, token)
		c.JSON(http.StatusOK, resource)
	}
}
