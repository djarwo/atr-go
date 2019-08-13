package atr

import (
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/helpers"
)

// Usecase is the contract between Repository and usecase
type Usecase interface {
	FindAll(helpers.FindAllParams) ([]*models.Atr, int, int, error)
}
