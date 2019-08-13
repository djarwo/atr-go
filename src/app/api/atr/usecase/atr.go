package usecase

import (
	"time"

	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/app/api/atr"
	"github.com/atomic/atr/src/helpers"
	"github.com/spf13/viper"
)

//AtrUsecase initialize object from model Atr, to be used in database operation
type AtrUsecase struct {
	atrRepo        atr.Repository
	contextTimeout time.Duration
}

//NewAtrUsecase initialize object from model Atr, to be used in database operation
func NewAtrUsecase(db *models.DB, atrRepo atr.Repository) atr.Usecase {
	//p := repository.NewAtrRepository(db)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	return &AtrUsecase{
		atrRepo:        atrRepo,
		contextTimeout: timeoutContext,
	}
}

//FindAll initialize object from model Atr, to be used in database operation
func (u *AtrUsecase) FindAll(filterFindAllParams helpers.FindAllParams) ([]*models.Atr, int, int, error) {

	result, length, statusCode, err := u.atrRepo.FindAll(filterFindAllParams)

	return result, length, statusCode, err
}
