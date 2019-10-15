package routes

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	// "github.com/atomic/atr/middleware"
	"github.com/atomic/atr/models"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes is a base function to register all routes (api and web)
func RegisterRoutes(db *models.DB) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowBrowserExtensions: true,
		// AllowAllOrigins:        true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Accept", "Accept-Language", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	RegisterAPIRoutes(db, router)

	envname := "PORTATOMICGO"
	port := os.Getenv(envname)
	serverAddress := port
	router.Run(serverAddress)

}
