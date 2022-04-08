package license_get

import (
	"github.com/gin-gonic/gin"
	"licenses/data/model"
	"licenses/data/repository"
	"net/http"
)

func Handler(repository *repository.LicenseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		uid, e := c.Get("uid")
		if !e {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to find id token."})
			return
		}

		license := &model.License{}

		err := repository.Get(id, license)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Couldn't find that license."})
			return
		}

		if license.Creator == "" || license.Creator != uid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}

		c.JSON(http.StatusOK, license)
	}
}
