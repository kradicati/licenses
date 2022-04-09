package license

import (
	"github.com/gin-gonic/gin"
	"licenses/data/repository"
	"net/http"
)

func Delete(repository *repository.LicenseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := repository.Delete(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}
