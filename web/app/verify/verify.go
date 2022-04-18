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

		now := time.Now().Unix()
		success := true
		reason := ""

		err := repository.Get(id, license)
		if err != nil {
			success = false
			reason = err.Error()
			//TODO log this as well
			internal.Abort(c, http.StatusNotFound)
			return
		}

		if success && license.Expires != 0 && license.Expires-now < 0 {
			success = false
			reason = "expired"
		}

		ip := c.ClientIP()

		if success && license.WhitelistedIps != nil && len(license.WhitelistedIps) > 0 && !internal.StringInSlice(ip, license.WhitelistedIps) {
			success = false
			reason = "not whitelisted"
		}

		if success && license.BlacklistedIps != nil && len(license.BlacklistedIps) > 0 && internal.StringInSlice(ip, license.BlacklistedIps) {
			success = false
			reason = "blacklisted"
		}

		ipLog := model.IpLog{
			Ip:     ip,
			Time:   now,
			Status: success,
			Reason: reason,
		}

		err = repository.UpdateIpLog(license.Id, append(license.IpLog, ipLog)...)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert ip into log."})
			return
		}

		if !success {
			internal.Abort(c, http.StatusTeapot)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
	}
}
