package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Adding CORS Header for microservice
func CORSHeader(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		// AllowBrowserExtensions: true,
		// AllowAllOrigins:        true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Accept", "Accept-Language", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
