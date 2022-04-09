package middleware

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	fbToken             = "FIREBASE_ID_TOKEN"
)

// https://github.com/MousyBusiness/AppEngineAPI/blob/master/internal/mauth/middleware.go

// AuthJWT Gin middleware for JWT auth
func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authenticated, e := c.Get("authenticated")
		if e && authenticated == "apiKey" {
			c.Next()
			return
		}

		authHeader := c.Request.Header.Get(authorizationHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := client.VerifyIDToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		c.Set("token", idToken)
		c.Set("uid", idToken.UID)
		c.Set("authenticated", "jwt")
		c.Next()
	}
}

/*
// AuthAPIKey API key auth middleware
func AuthAPIKey(secretId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get(apiKeyHeader)

		secret, err := secrets.GetSecret(secretId)
		if err != nil {
			log.Println("failed to get secret")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("comparing secret with provided key", secret, key)

		if secret != key {
			log.Println("key doesnt match!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("no error during check")
		c.Next()
	}
}
*/
