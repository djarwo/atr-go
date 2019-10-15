package http_user_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/atomic/atr/library/faker"
	"github.com/atomic/atr/models"
	UserHttp "github.com/atomic/atr/src/app/api/user/delivery/http_user"
	"github.com/atomic/atr/src/app/api/user/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFindAll(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)

	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	mockListUser := make([]*models.User, 0)
	mockListUser = append(mockListUser, &mockUser)

	fmt.Println(mockListUser)
	mockUCase.On("FindAll", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListUser, nil)

	handler := UserHttp.UserHandler{
		UserUsecase: mockUCase,
	}

	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	handler.UserFindAll(c)
	require.NoError(t, err)
	mockUCase.AssertExpectations(t)

	/*r := library.GetRouter(true)
	r.GET("/api/v1/user/1", handler.UserFindAll)

	req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})*/

}
