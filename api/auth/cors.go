package auth

import cors "gopkg.in/gin-contrib/cors.v1"

func Cors(config *cors.Config) {
	config.AddExposeHeaders("Authorization")
}
