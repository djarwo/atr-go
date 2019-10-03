package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/app/api/user"
	"github.com/atomic/atr/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//UserUsecase initialize object from model User, to be used in database operation
type UserUsecase struct {
	userRepo       user.Repository
	contextTimeout time.Duration
}

//NewUserUsecase initialize object from model User, to be used in database operation
func NewUserUsecase(db *models.DB, userRepo user.Repository) user.Usecase {
	//p := repository.NewUserRepository(db)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	return &UserUsecase{
		userRepo:       userRepo,
		contextTimeout: timeoutContext,
	}
}

//Find initialize object from model User, to be used in database operation
func (u *UserUsecase) Login(username string, password string, deviceId string, notificationToken string) (*models.User, interface{}, int, error) {
	var err error
	if password == "" {
		err = fmt.Errorf("Password tidak boleh kosong")
		return nil, nil, http.StatusNotFound, err
	}

	if username == "" {
		err = fmt.Errorf("Username tidak boleh kosong")
		return nil, nil, http.StatusNotFound, err
	}

	return u.userRepo.Login(username, password, deviceId, notificationToken)
}

//FindAll initialize object from model User, to be used in database operation
func (u *UserUsecase) FindAll(filterFindAllParams helpers.FindAllParams) ([]*models.User, int, int, error) {

	result, length, statusCode, err := u.userRepo.FindAll(filterFindAllParams)

	return result, length, statusCode, err
}

//Find initialize object from model User, to be used in database operation
func (u *UserUsecase) Find(id string) (*models.User, int, error) {

	if id == "" {
		return nil, http.StatusNotFound, nil
	}

	return u.userRepo.Find(id)
}

//Create initialize object from model User, to be used in database operation
func (u *UserUsecase) Create(obj models.User, c *gin.Context) (*models.User, int, error) {
	return u.userRepo.Create(obj, c.PostForm("UserOutlet"), c.PostForm("UserCustomerID"))
}

//Update initialize object from model User, to be used in database operation
func (u *UserUsecase) Update(id string, obj models.User, c *gin.Context) (*models.User, int, error) {

	if id == "" {
		return nil, http.StatusNotFound, nil
	}

	return u.userRepo.Update(id, obj, c.PostForm("UserOutlet"), c.PostForm("UserCustomerID"))
}

//Delete initialize object from model User, to be used in database operation
func (u *UserUsecase) Delete(id string) (*models.User, int, error) {
	if id == "" {
		return nil, http.StatusNotFound, nil
	}

	return u.userRepo.Delete(id)
}