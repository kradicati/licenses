package middleware

import (
	"github.com/gin-gonic/gin"
	"licenses/internal"
)

func CheckRoles(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !internal.HasRole(c, role) {
			return
		}

		c.Next()
	}
}
