package routes

import (
	"github.com/atomic/atr/models"
	restapi "github.com/atomic/atr/src/app/api"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes is a base function to register all API Routes in the project
func RegisterAPIRoutes(db *models.DB, router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		restapi.RegisterRoutes(db, router, v1)
	}

}
