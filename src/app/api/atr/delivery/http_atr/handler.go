package http_atr

import (
	"net/http"

	"github.com/atomic/atr/middleware"

	"github.com/atomic/atr/library"

	"github.com/atomic/atr/models"

	"github.com/atomic/atr/src/app/api/atr"
	"github.com/atomic/atr/src/app/api/atr/repository"
	"github.com/atomic/atr/src/app/api/atr/usecase"
	"github.com/atomic/atr/src/helpers"
	"github.com/gin-gonic/gin"
)

var (
	response     interface{}
	findallparam helpers.FindAllParams
)

// AtrHandler  represent the httphandler for article
type AtrHandler struct {
	AtrUsecase atr.Usecase
	Result     gin.H
	Status     int
}

// RegisterAPI  represent the httphandler for article
func (h AtrHandler) RegisterAPI(db *models.DB, router *gin.Engine, v *gin.RouterGroup) {

	atrRepo := repository.NewAtrRepository(db)

	u := usecase.NewAtrUsecase(db, &atrRepo)

	base := &AtrHandler{AtrUsecase: u}

	rs := v.Group("/atrs")
	{
		rs.GET("", middleware.Auth, base.AtrFindAll)
	}

}

//AtrFindAll get all data in atr
func (h *AtrHandler) AtrFindAll(c *gin.Context) {
	page, size := helpers.FilterFindAll(c)
	claims := library.GetJWTClaims("")
	filterFindAllParams := helpers.FilterFindAllParam(c, claims)
	datas, length, statusCode, err := h.AtrUsecase.FindAll(filterFindAllParams)

	if length == 0 || err != nil {
		response = helpers.ResultAll{Status: "Warning", StatusCode: statusCode, Message: helpers.MessageErr(err)}
		h.Status = statusCode
		h.Result = gin.H{
			"result": response,
		}
		helpers.ReturnHandler(h.Status, err)
	} else {
		response = helpers.ResultAll{Status: "Sukses", StatusCode: http.StatusOK, Message: "Data Atr Berhasil Di Tampilkan", TotalData: length, Page: page, Size: size, Data: datas}
		h.Status = http.StatusOK
		h.Result = gin.H{
			"result": response,
		}
	}

	helpers.ReturnHandler(h.Status, err)
	c.JSON(h.Status, h.Result)
}
