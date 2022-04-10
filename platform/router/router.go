package router

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"licenses/data/repository"
	"licenses/platform/middleware"
	"licenses/web/app/licenses"
	"licenses/web/app/licenses/license"
	"licenses/web/app/users"
	"licenses/web/app/verify"
)

var (
	licenseRepository *repository.LicenseRepository
	userRepository    *repository.UserRepository
)

func New(client *auth.Client, store *firestore.Client) *gin.Engine {
	router := gin.Default()

	initFireStore(store)

	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:*"}, // TODO FIX
			AllowMethods:     []string{"PUT", "PATCH"},
			AllowHeaders:     []string{"Origin Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				//return origin == "*://localhost:8008"
				return true
			},
			MaxAge: 12 * time.Hour,
		}))
	*/

	router.Use(configCors())

	router.GET("/api/v1/verify/:id", verify.Handler(licenseRepository))

	{ // v1
		v1 := router.Group("/api/v1")

		v1.Use(middleware.AuthJWT(client))
		{ // licenses
			licenseGroup := v1.Group("/licenses")

			licenseGroup.Use(middleware.CheckAPIKey(userRepository))

			licenseGroup.POST("/", licenses.Create(licenseRepository, userRepository))
			licenseGroup.GET("/", licenses.List(licenseRepository))

			{ // :id
				mgmt := licenseGroup.Group("/:id")

				mgmt.Use(middleware.CheckLicenseAccess(licenseRepository))

				mgmt.GET("/", license.Get)
				mgmt.DELETE("/", license.Delete(licenseRepository))
			}
		}

		{ // users
			userGroup := v1.Group("/users")

			userGroup.POST("/", middleware.CheckRoles("admin"), users.Create(client))
		}
	}

	return router
}

func configCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "X-Requested-With", "Accept", "Authorization")
	return cors.New(config)
}

func initFireStore(store *firestore.Client) {
	licenseRepository = repository.NewLicenseRepository(store)
	userRepository = repository.NewUserRepository(store)
}
