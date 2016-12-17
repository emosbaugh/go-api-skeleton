package v1

import (
	"database/sql"
	"net/http"

	"github.com/replicatedcom/gin-example/api/context"
	"github.com/replicatedcom/gin-example/api/errors"
	"github.com/replicatedcom/gin-example/api/routes/v1/resource"
	"github.com/replicatedcom/gin-example/api/routes/v1/schema"
	"github.com/replicatedcom/gin-example/api/schemas"
	"github.com/replicatedcom/gin-example/inject"

	log "github.com/Sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/gin-gonic/gin.v1"
)

var userResourceSchema *gojsonschema.Schema

func init() {
	userResourceSchema = schema.MustLoadSchema("file:///user_resource.json")
}

// UserRetrieve retrieves the current user from the context.
func UserRetrieve(env *inject.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := context.GetUser(c, env)
		if err != nil {
			if err == sql.ErrNoRows {
				errors.ErrResourceNotFound.Response(c)
				return
			}
			log.WithField("err", err).Error("failed to retrieve user")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if user == nil {
			log.Error("user not found in context")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		resource := resource.NewUser(user)
		err = schemas.MaybeValidateResponse(c, resource, userResourceSchema)
		if err != nil {
			log.WithField("err", err).Error("failed to validate resource")
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}
