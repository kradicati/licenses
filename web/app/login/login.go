package login

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	username, password string
}

func Handle(client *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("yes")

		var login Login

		err := ctx.Bind(&login)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": http.StatusText(http.StatusBadRequest),
			})
		}

		log.Println(login)

		/*
			client.GetUserByEmail()

			claims := map[string]interface{}{
				"roles": []string{"user"},
			}

			token, err := client.CustomTokenWithClaims(ctx, "some-uid", claims)
			if err != nil {
				log.Fatalf("error minting custom token: %v\n", err)
			}

		*/
	}
}
