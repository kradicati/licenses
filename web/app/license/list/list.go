package list

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"licenses/data/model"
	"licenses/data/repository"
	"net/http"
)

func Handler(repository *repository.LicenseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, e := c.Get("uid")
		if !e {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to find id token."})
			return
		}

		licenses := &[]model.License{}

		err := repository.FindAllByUser(uid.(string), licenses)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, *licenses)
	}
}
