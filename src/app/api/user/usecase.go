package user

import (
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/helpers"
	"github.com/gin-gonic/gin"
)

// Usecase is the contract between Repository and usecase
type Usecase interface {
	FindAll(helpers.FindAllParams) ([]*models.User, int, int, error)
	Find(string) (*models.User, int, error)
	Create(models.User, *gin.Context) (*models.User, int, error)
	Update(string, models.User, *gin.Context) (*models.User, int, error)
	Delete(string) (*models.User, int, error)
	Login(string, string, string, string) (*models.User, interface{}, int, error)
}
