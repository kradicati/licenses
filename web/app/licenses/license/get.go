package license

import (
	"github.com/gin-gonic/gin"
	"licenses/internal"
	"net/http"
)

func Get(c *gin.Context) {
	license, e := c.Get("licenses")

	if !e {
		internal.Abort(c, http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, license)
}
