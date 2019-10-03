package repository

import (
	"net/http"
	"strconv"
	"github.com/sirupsen/logrus"
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/helpers"

)

//AtrRepository initialize object from model Atr, to be used in database operation
type AtrRepository struct {
	objs       []*models.Atr
	obj        models.Atr
	objsLength []models.Atr
	db         *models.DB
	r          helpers.ReturnRepo
}

//NewAtrRepository initialize service that provide connection to Database
func NewAtrRepository(db *models.DB) AtrRepository {
	//db := &models.DB{DB: configs.ActiveDB}
	return AtrRepository{db: db, r: helpers.DefaultRepo()}
}

//FindAll is a function to get all Data
func (s AtrRepository) FindAll(params helpers.FindAllParams) ([]*models.Atr, int, int, error) {
	var err error
	var length int

	db := s.db.DB.Set("gorm:auto_preload", true)
	if params.Page != "" && params.Page != "-1" {
		countpage, _ := strconv.Atoi(params.Page)
		countsize, _ := strconv.Atoi(params.Size)
		startdata := (countpage * countsize) - countsize

		if params.Keyword != "" {
			err = db.Offset(startdata).Limit(countsize).Find(&s.objs).Error
			s.db.DB.Find(&s.objsLength)
			length = len(s.objsLength)
		} else {
			err = db.Offset(startdata).Limit(countsize).Find(&s.objs).Error
			s.db.DB.Find(&s.objsLength)
			length = len(s.objsLength)
		}

	} else {
		err = db.Find(&s.objs).Error
		length = len(s.objs)
	}

	if err != nil {
		logrus.Error(err)
		s.r.StatusCode = http.StatusInternalServerError
		return nil, 0, s.r.StatusCode, err
	}
	return s.objs, length, s.r.StatusCode, nil
}
