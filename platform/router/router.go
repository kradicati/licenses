package router

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"licenses/platform/middleware/fbauth"
)

func New(client *auth.Client) *gin.Engine {
	router := gin.Default()

	router.Use(fbauth.AuthJWT(client))

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Success!"})
	})

	return router
}
