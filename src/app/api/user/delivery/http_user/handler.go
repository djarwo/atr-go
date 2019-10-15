package http_user

import (
	"net/http"

	"github.com/atomic/atr/models"

	"github.com/atomic/atr/library"

	"github.com/atomic/atr/src/helpers"

	"github.com/atomic/atr/middleware"
	"github.com/atomic/atr/src/app/api/user"
	"github.com/atomic/atr/src/app/api/user/repository"
	"github.com/atomic/atr/src/app/api/user/usecase"
	"github.com/gin-gonic/gin"
)

var (
	response     interface{}
	findallparam helpers.FindAllParams
)

// UserHandler  represent the httphandler for article
type UserHandler struct {
	UserUsecase user.Usecase
	Result      gin.H
	Status      int
}

// RegisterAPI  represent the httphandler for article
func (h UserHandler) RegisterAPI(db *models.DB, router *gin.Engine, v *gin.RouterGroup) {

	userRepo := repository.NewUserRepository(db)

	u := usecase.NewUserUsecase(db, &userRepo)

	base := &UserHandler{UserUsecase: u}

	rs := v.Group("/users")
	{
		rs.GET("", middleware.Auth, base.UserFindAll)
	}

	auth := v.Group("/auth")
	{
		auth.POST("/user/login", base.UserLogin)
	}

	r := v.Group("/user")
	{
		r.GET("/:id", middleware.Auth, base.UserFind)

		r.POST("", middleware.Auth, base.UserCreate)
		r.PUT("/:id", middleware.Auth, base.UserUpdate)
		r.DELETE("/:id", middleware.Auth, base.UserDelete)
	}
}

func (h *UserHandler) UserLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	deviceId := c.PostForm("device_token")
	notificationToken := c.PostForm("notification_token")

	data, param, statusCode, err := h.UserUsecase.Login(username, password, deviceId, notificationToken)

	if err != nil {
		response = helpers.Result{Status: "Warning", StatusCode: statusCode, Message: helpers.MessageErr(err), Data: param}
		h.Status = statusCode
		h.Result = gin.H{
			"result": response,
		}
		helpers.ReturnHandler(h.Status, err)
	} else {
		response = helpers.Result{Status: "Sukses", StatusCode: http.StatusOK, Message: "Login Berhasil", Data: data}
		h.Status = http.StatusOK

		h.Result = gin.H{
			"result": response,
		}
	}

	c.JSON(h.Status, h.Result)

}

//UserFindAll get all data in User
func (h *UserHandler) UserFindAll(c *gin.Context) {
	page, size := helpers.FilterFindAll(c)
	claims := library.GetJWTClaims("")
	filterFindAllParams := helpers.FilterFindAllParam(c, claims)

	datas, length, statusCode, err := h.UserUsecase.FindAll(filterFindAllParams)
	if length == 0 || err != nil {
		response = helpers.ResultAll{Status: "Warning", StatusCode: statusCode, Message: helpers.MessageErr(err)}
		h.Status = statusCode
		h.Result = gin.H{
			"result": response,
		}
		helpers.ReturnHandler(h.Status, err)
	} else {
		response = helpers.ResultAll{Status: "Sukses", StatusCode: http.StatusOK, Message: "Data User Berhasil Di Tampilkan", TotalData: length, Page: page, Size: size, Data: datas}
		h.Status = http.StatusOK
		h.Result = gin.H{
			"result": response,
		}
	}

	c.JSON(h.Status, h.Result)
}

//UserFind get data in user by id
func (h *UserHandler) UserFind(c *gin.Context) {

	id := c.Param("id")

	data, statusCode, err := h.UserUsecase.Find(id)

	if data == nil || err != nil {
		response = helpers.Result{Status: "Warning", StatusCode: statusCode, Message: helpers.MessageErr(err)}
		h.Status = statusCode
		h.Result = gin.H{
			"result": response,
		}
		helpers.ReturnHandler(h.Status, err)
	} else {
		response = helpers.Result{Status: "Sukses", StatusCode: http.StatusOK, Message: "Data User Berhasil Ditampilkan", Data: data}
		h.Status = http.StatusOK
		h.Result = gin.H{
			"result": response,
		}
	}

	c.JSON(h.Status, h.Result)
}

//UserCreate create user data
func (h *UserHandler) UserCreate(c *gin.Context) {
	var user models.User

	password := library.PasswordHasher(c.PostForm("Password"))
	user.Password = password
	user.Code = helpers.CodeGenerator("User", "USR")
	user.Username = c.PostForm("Username")
	user.Name = c.PostForm("Name")
	user.Email = c.PostForm("Email")
	user.Phone = c.PostForm("Phone")
	user.Pin = c.PostForm("Pin")
	user.LoginType = c.PostForm("LoginType")
	user.UID = c.PostForm("UID")
	user.Description = c.PostForm("Description")

	data, statusCode, err := h.UserUsecase.Create(user, c)

	if data == nil || err != nil {
		response = helpers.Result{Status: "Warning", StatusCode: statusCode, Message: helpers.MessageErr(err)}
		h.Status = statusCode
		h.Result = gin.H{
			"result": response,
		}
		helpers.ReturnHandler(h.Status, err)
	} else {
		response = helpers.Result{Status: "Sukses", StatusCode: http.StatusOK, Message: "Data User Berhasil Ditambahkan", Data: data}
		h.Status = http.StatusOK
		h.Result = gin.H{
			"result": response,
		}
	}

	c.JSON(h.Status, h.Result)
}

//UserUpdate update data in user by id
func (h *UserHandler) UserUpdate(c *gin.Context) {
	var user models.User

	id := c.Param("id")

	password := library.PasswordHasher(c.PostForm("Password"))
	user.Username = c.PostForm("Username")
	user.Password = password
	user.Name = c.PostForm("Name")
	user.Email = c.PostForm("Email")
	user.Phone = c.PostForm("Phone")
	user.Pin = c.PostForm("Pin")
	user.LoginType = c.PostForm("LoginType")
	user.UID = c.PostForm("UID")
	user.Description = c.PostForm("Description")
	data, statusCode, err := h.UserUsecase.Update(id, user, c)

	if data == nil || err != nil {
		response = helpers.Result{Status: "Warning", StatusCode: statusCode, Message: helpers.MessageErr(err)}
		h.Status = statusCode
		h.Result = gin.H{
			"result": response,
		}
		helpers.ReturnHandler(h.Status, err)
	} else {
		response = helpers.Result{Status: "Sukses", StatusCode: http.StatusOK, Message: "Data User Berhasil Diubah", Data: data}
		h.Status = http.StatusOK
		h.Result = gin.H{
			"result": response,
		}
	}

	c.JSON(h.Status, h.Result)
}

//UserDelete delete data in user by id
func (h *UserHandler) UserDelete(c *gin.Context) {

	id := c.Param("id")

	data, statusCode, err := h.UserUsecase.Delete(id)

	if data == nil || err != nil {
		response = helpers.Result{Status: "Warning", StatusCode: statusCode, Message: helpers.MessageErr(err)}
		h.Status = statusCode
		h.Result = gin.H{
			"result": response,
		}
		helpers.ReturnHandler(h.Status, err)
	} else {
		response = helpers.Result{Status: "Sukses", StatusCode: http.StatusOK, Message: "Data User Berhasil Dihapus", Data: data}
		h.Status = http.StatusOK
		h.Result = gin.H{
			"result": response,
		}
	}

	c.JSON(h.Status, h.Result)
}
