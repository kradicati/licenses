package internal

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringInInterfaces(a string, list []interface{}) bool {
	for _, b := range list {
		if b.(string) == a {
			return true
		}
	}
	return false
}

func Abort(c *gin.Context, status int) {
	c.AbortWithStatusJSON(status, gin.H{"message": http.StatusText(status)})
}

func HasRole(c *gin.Context, role string) bool {
	token, e := c.Get("token")
	if !e {
		Abort(c, http.StatusNotFound)
		return false
	}

	claims := token.(*auth.Token).Claims

	roles, ok := claims["roles"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No roles claim."})
		return false
	}

	if !StringInInterfaces(role, roles.([]interface{})) {
		Abort(c, http.StatusUnauthorized)
		return false
	}

	return true
}
