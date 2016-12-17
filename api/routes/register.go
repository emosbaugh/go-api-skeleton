package routes

import (
	"github.com/replicatedcom/gin-example/api/routes/v1"
	"github.com/replicatedcom/gin-example/api/routes/v1/schema"
	"github.com/replicatedcom/gin-example/inject"

	"gopkg.in/gin-gonic/gin.v1"
)

func Register(env *inject.Env, public, private *gin.RouterGroup) {
	RegisterV1(env, public.Group("/v1"), private.Group("/v1"))
}

func RegisterV1(env *inject.Env, public, private *gin.RouterGroup) {
	public.StaticFS("/schema", schema.AssetFS)

	public.POST("/user", v1.UserCreate(env))
	public.POST("/login", v1.UserLogin(env))

	private.GET("/refresh-token", v1.RefreshToken)
	private.GET("/user", v1.UserRetrieve(env))
}
