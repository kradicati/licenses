package router

import (
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"licenses/platform/middleware/fbauth"
	"licenses/web/app/login"
	"time"
)

func New(client *auth.Client) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // TODO FIX
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "*://localhost:8008"
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	v1.Use(fbauth.AuthJWT(client))

	v1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Authenticated"})
	})

	router.POST("/login", login.Handle(client))

	return router
}
