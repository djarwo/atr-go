package helpers

import (
	"net/http"
	"github.com/sirupsen/logrus"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {

	default:
		return http.StatusInternalServerError
	}
}
