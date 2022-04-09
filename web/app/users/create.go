package users

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userCredentials struct {
	Email       string   `json:"email,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	Password    string   `json:"password,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}

func Create(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &userCredentials{}

		err := c.BindJSON(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
			return
		}

		params := (&auth.UserToCreate{}).
			Email(user.Email).
			EmailVerified(true).
			Password(user.Password).
			DisplayName(user.DisplayName).
			Disabled(false)

		u, err := authClient.CreateUser(context.Background(), params)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError) + ": " + err.Error()})
			return
		}

		err = authClient.SetCustomUserClaims(context.Background(), u.UID, map[string]interface{}{"roles": user.Roles})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError) + ": " + err.Error()})
			return
		}
	}
}
