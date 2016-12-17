package v1

import (
	"net/http"

	"github.com/replicatedcom/gin-example/api/auth"
	"github.com/replicatedcom/gin-example/api/routes/v1/params"
	"github.com/replicatedcom/gin-example/api/routes/v1/resource"
	"github.com/replicatedcom/gin-example/api/routes/v1/schema"
	"github.com/replicatedcom/gin-example/api/schemas"
	"github.com/replicatedcom/gin-example/inject"
	"github.com/replicatedcom/gin-example/models"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/gin-gonic/gin.v1"
)

var userCreateParamsSchema *gojsonschema.Schema

func init() {
	userCreateParamsSchema = schema.MustLoadSchema("file:///user_create_params.json")
}

func UserCreate(env *inject.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		schemas.LinkHeader(c, "schema.json#/properties/userCreateParams")

		var request params.UserCreate
		err := schemas.Bind(c, &request, userCreateParamsSchema)
		if err != nil {
			log.WithField("err", err).Error("failed to bind to login params")
			return
		}

		user := &models.User{
			Email: request.Email,
		}
		err = env.UserService.Create(user, request.Password)
		if err != nil {
			// TODO: error for duplicate
			log.WithField("err", err).Error("failed to create user")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

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
		c.JSON(http.StatusCreated, resource)
	}
}
