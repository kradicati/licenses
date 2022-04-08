package create

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"licenses/data/model"
	"licenses/data/repository"
	"net/http"
	"time"
)

func Handler(repository *repository.LicenseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		license := &model.License{}

		err := c.BindJSON(license)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to bind body. " + err.Error()})
			return
		}

		uid, e := c.Get("uid")
		if !e {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to find id token."})
			return
		}

		license.Creator = uid.(string)
		license.Created = time.Now().Unix()
		license.IpLog = make([]string, 0)
		license.WhitelistedIps = makeIfAbsent(license.WhitelistedIps)
		license.BlacklistedIps = makeIfAbsent(license.BlacklistedIps)

		fmt.Println()

		err = repository.Insert(license.Id, license)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert license into database." + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Success."})
	}
}

func makeIfAbsent(arr []string) []string {
	if arr == nil {
		arr = []string{}
	}

	return arr
}
