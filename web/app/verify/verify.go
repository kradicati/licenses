package verify

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"licenses/data/model"
	"licenses/data/repository"
	"licenses/internal"
	"net/http"
	"time"
)

func Handler(repository *repository.LicenseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO rate-limit, cache, asymmetric key encryption(?)

		id := c.Param("id")

		license := &model.License{}

		err := repository.Get(id, license)
		if err != nil {
			abort(c, err)
			return
		}

		if license.Expires != 0 && license.Expires-time.Now().Unix() < 0 {
			abort(c, err)
			return
		}

		ip := c.ClientIP()

		if license.WhitelistedIps != nil && len(license.WhitelistedIps) > 0 && !internal.StringInSlice(ip, license.WhitelistedIps) {
			abort(c, err)
			return
		}

		if license.BlacklistedIps != nil && len(license.BlacklistedIps) > 0 && internal.StringInSlice(ip, license.BlacklistedIps) {
			abort(c, err)
			return
		}

		err = repository.UpdateIpLog(license.Id, append(license.IpLog, ip)...)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert ip into log."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}

func abort(c *gin.Context, err ...error) {
	c.AbortWithStatusJSON(http.StatusTeapot, gin.H{"message": http.StatusText(http.StatusTeapot)})
}
