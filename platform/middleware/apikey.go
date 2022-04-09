package middleware

import (
	"github.com/gin-gonic/gin"
	"licenses/data/model"
	"licenses/data/repository"
	"net/http"
)

func CheckAPIKey(repository *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO Cache

		apiKey := c.Request.Header.Get("X-API-Key")

		if apiKey == "" {
			c.Next()
			return
		}

		users := &[]model.User{}

		err := repository.FindAllByApiKey(apiKey, users)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
			return
		}

		if len(*users) != 1 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
			return
		}

		user := (*users)[0]

		c.Set("uid", user.Id)
		c.Set("authenticated", "apiKey")
		c.Next()
	}
}
