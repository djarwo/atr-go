package example

import (
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/app/api/atr/delivery/http_atr"
	"github.com/atomic/atr/src/app/api/user/delivery/http_user"
	"github.com/gin-gonic/gin"
)

var (
	atrHandler  http_atr.AtrHandler
	userHandler http_user.UserHandler
)

func RegisterRoutes(db *models.DB, router *gin.Engine, v *gin.RouterGroup) {
	atrHandler.RegisterAPI(db, router, v)
	userHandler.RegisterAPI(db, router, v)
}
