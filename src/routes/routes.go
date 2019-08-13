package routes

import (
	"github.com/atomic/atr/middleware"
	"github.com/atomic/atr/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// RegisterRoutes is a base function to register all routes (api and web)
func RegisterRoutes(db *models.DB) {
	router := gin.Default()

	middleware.CORSHeader(router)

	RegisterAPIRoutes(db, router)

	serverAddress := viper.GetString(`server.address`)
	router.Run(serverAddress)

}
