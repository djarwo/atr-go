package atr

import (
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/helpers"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(helpers.FindAllParams) ([]*models.Atr, int, int, error)
}
