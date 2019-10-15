package repository

import (
	"net/http"
	"strconv"

	"github.com/atomic/atr/library"
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//UserRepository initialize object from model User, to be used in database operation
type UserRepository struct {
	objs       []*models.User
	obj        models.User
	objsLength []models.User
	db         *models.DB
	r          helpers.ReturnRepo
}

//NewUserRepository initialize service that provide connection to Database
func NewUserRepository(db *models.DB) UserRepository {
	//db := &models.DB{DB: configs.ActiveDB}
	return UserRepository{db: db, r: helpers.DefaultRepo()}
}

func (s UserRepository) Login(username string, password string, deviceToken string, notificationToken string) (*models.User, interface{}, int, error) {

	pwd := library.PasswordHasher(password)

	err := s.db.DB.Set("gorm:auto_preload", true).Where(`(email=? OR phone=? OR username=?) AND password=?`, username, username, username, pwd).Find(&s.obj).Error

	if err != nil {
		params := gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password, with error message: " + err.Error(),
		}
		s.r.StatusCode = http.StatusUnauthorized
		return nil, params, s.r.StatusCode, err
	} else {
		credentials := helpers.Credential{ID: s.obj.ID, Username: s.obj.Username, Email: s.obj.Email}
		token, err := library.JwtSignString(credentials)
		if err != nil {
			params := gin.H{
				"status":  http.StatusInternalServerError,
				"message": "internal server error",
			}
			s.r.StatusCode = http.StatusInternalServerError
			return nil, params, s.r.StatusCode, err
		} else {
			params := gin.H{
				"status":  http.StatusOK,
				"message": s.obj.Email + " loggedin",
				"token":   token,
			}

			var userToken models.User
			userToken.DeviceToken = deviceToken
			userToken.NotificationToken = notificationToken
			userToken.SessionToken = token

			s.db.DB.Model(&s.obj).Updates(userToken)

			return &s.obj, params, s.r.StatusCode, nil
		}
	}
}

//FindAll is a function to get all Data
func (s UserRepository) FindAll(params helpers.FindAllParams) ([]*models.User, int, int, error) {
	var err error
	var length int

	db := s.db.DB.Set("gorm:auto_preload", true)
	if params.Page != "" && params.Page != "-1" {
		countpage, _ := strconv.Atoi(params.Page)
		countsize, _ := strconv.Atoi(params.Size)
		startdata := (countpage * countsize) - countsize

		if params.Keyword != "" {
			err = db.Offset(startdata).Limit(countsize).Where(params.Query).Where(params.DataFinder).Order(params.SortBy).Group(params.GroupBy).Find(&s.objs).Error
			s.db.DB.Where(params.Query).Where(params.DataFinder).Where(params.StatusID + params.BusinessID).Group(params.GroupBy).Find(&s.objsLength)
			length = len(s.objsLength)
		} else {
			err = db.Offset(startdata).Limit(countsize).Where(params.Query).Where(params.DataFinder).Where(params.StatusID + params.BusinessID).Order(params.SortBy).Group(params.GroupBy).Find(&s.objs).Error
			s.db.DB.Where(params.Query).Where(params.DataFinder).Where(params.StatusID + params.BusinessID).Group(params.GroupBy).Find(&s.objsLength)
			length = len(s.objsLength)
		}

	} else {
		err = db.Where(params.Query).Where(params.DataFinder).Where(params.StatusID + params.BusinessID).Order(params.SortBy).Find(&s.objs).Error
		length = len(s.objs)
	}

	if err != nil {
		logrus.Error(err)
		s.r.StatusCode = http.StatusNotFound
		return nil, 0, s.r.StatusCode, err
	}

	return s.objs, length, s.r.StatusCode, nil
}

//Find is a function to get by ID
func (s UserRepository) Find(id string) (*models.User, int, error) {

	err := s.db.DB.Set("gorm:auto_preload", true).Where(library.GetBusiness()).Where("id = ?", id).First(&s.obj).Error

	if err != nil {
		logrus.Error(err)
		s.r.StatusCode = http.StatusInternalServerError
		return nil, s.r.StatusCode, err
	}

	return &s.obj, s.r.StatusCode, nil
}

//Create is a function to get by ID
func (s UserRepository) Create(obj models.User) (*models.User, int, error) {
	err := s.db.DB.Create(&obj).Error

	if err != nil {
		logrus.Error(err)
		s.r.StatusCode = http.StatusInternalServerError
		return nil, s.r.StatusCode, err
	}

	var user models.User
	s.db.DB.Find(&user, obj.ID)
	return &user, s.r.StatusCode, nil
}

//Update is a function to get by ID
func (s UserRepository) Update(id string, obj models.User) (*models.User, int, error) {
	err := s.db.DB.Set("gorm:auto_preload", true).First(&s.obj, id).Error
	if err != nil {
		s.r.StatusCode = http.StatusInternalServerError
		return nil, s.r.StatusCode, err
	}

	errUpdate := s.db.DB.Model(&s.obj).Updates(obj).Error

	if errUpdate != nil {
		logrus.Error(errUpdate)
		s.r.StatusCode = http.StatusInternalServerError
		return nil, s.r.StatusCode, errUpdate
	}

	var user models.User
	s.db.DB.Find(&user, s.obj.ID)
	return &user, s.r.StatusCode, nil
}

//Delete is a function to get by ID
func (s UserRepository) Delete(id string) (*models.User, int, error) {
	err := s.db.DB.Where("id = ?", id).Delete(&s.obj).Error

	if err != nil {
		logrus.Error(err)
		s.r.StatusCode = http.StatusInternalServerError
		return nil, s.r.StatusCode, err
	}

	return &s.obj, s.r.StatusCode, nil
}
