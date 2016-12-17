package api

import (
	"github.com/replicatedcom/gin-example/api/auth"
	"github.com/replicatedcom/gin-example/api/context"
	"github.com/replicatedcom/gin-example/api/errors"
	"github.com/replicatedcom/gin-example/api/routes"
	"github.com/replicatedcom/gin-example/inject"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gin-contrib/cors.v1"
	"gopkg.in/gin-gonic/gin.v1"
)

func Run(env *inject.Env) error {
	if log.GetLevel() != log.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(log.StandardLogger().Writer()),
		gin.Recovery(),
		errors.Middleware(),
	)

	// json 404
	r.NoRoute(func(c *gin.Context) {
		errors.ErrNotFound.Response(c)
	})

	// cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	auth.Cors(&config)
	r.Use(cors.New(config))

	// auth
	public := r.Group("/api")
	private := public.Group(
		"",
		auth.Middleware(viper.GetString("secret"), log.StandardLogger()),
		context.Middleware(log.StandardLogger()),
	)
	routes.Register(env, public, private)

	// Listen and server on 0.0.0.0:8080
	return r.Run(":8080")

	// TODO: binding json only + csrf
	// TODO: invalidate sessions on password change
	// TODO: api token auth
	// TODO: tests
}
