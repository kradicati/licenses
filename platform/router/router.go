package router

import (
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"licenses/platform/middleware/fbauth"
	"licenses/web/app/login"
)

func New(client *auth.Client) *gin.Engine {
	router := gin.Default()

	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:*"}, // TODO FIX
			AllowMethods:     []string{"PUT", "PATCH"},
			AllowHeaders:     []string{"Origin Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				//return origin == "*://localhost:8008"
				return true
			},
			MaxAge: 12 * time.Hour,
		}))
	*/

	router.Use(configCors())

	v1 := router.Group("/api/v1")

	v1.Use(fbauth.AuthJWT(client))

	v1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Authenticated"})
	})

	router.POST("/login", login.Handle(client))

	return router
}

func configCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	return cors.New(config)
}
