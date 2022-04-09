package license

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	license, e := c.Get("licenses")

	if !e {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
		return
	}

	c.JSON(http.StatusOK, license)
}
