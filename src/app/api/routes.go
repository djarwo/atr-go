package example

import (
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/app/api/atr/delivery/http_atr"
	"github.com/gin-gonic/gin"
)

var (
	atrHandler http_atr.AtrHandler
)

func RegisterRoutes(db *models.DB, router *gin.Engine, v *gin.RouterGroup) {
	atrHandler.RegisterAPI(db, router, v)
}
