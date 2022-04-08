package create

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"licenses/data/model"
	"licenses/platform/middleware/fbauth"
	"net/http"
)

func Handler(c *gin.Context) {
	var license model.License

	err := c.Bind(license)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to bind body."})
		return
	}

	value, e := c.Get(fbauth.FbToken)
	if !e {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to find id token."})
		return
	}

	uid := value.(auth.Token).UID

	license.Creator = uid

	c.JSON(200, license)
}
