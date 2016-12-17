package context

import (
	"github.com/replicatedcom/gin-example/inject"
	"github.com/replicatedcom/gin-example/models"

	"gopkg.in/gin-gonic/gin.v1"
)

func SetUserID(c *gin.Context, userID string) {
	c.Set("user.id", userID)
}

func GetUser(c *gin.Context, env *inject.Env) (*models.User, error) {
	val, ok := c.Get("user.id")
	if !ok {
		return nil, nil
	}
	userID, ok := val.(string)
	if !ok {
		return nil, nil
	}
	return env.UserService.Get(userID)
}
