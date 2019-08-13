package http_atr_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/atomic/atr/library/faker"
	"github.com/atomic/atr/models"
	AtrHttp "github.com/atomic/atr/src/app/api/atr/delivery/http_atr"
	"github.com/atomic/atr/src/app/api/atr/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFindAll(t *testing.T) {

	var mockAtr models.Atr
	err := faker.FakeData(&mockAtr)

	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	mockListAtr := make([]*models.Atr, 0)
	mockListAtr = append(mockListAtr, &mockAtr)

	fmt.Println(mockListAtr)
	mockUCase.On("FindAll", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockListAtr, nil)

	handler := AtrHttp.AtrHandler{
		AtrUsecase: mockUCase,
	}

	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	handler.AtrFindAll(c)
	require.NoError(t, err)
	mockUCase.AssertExpectations(t)

	/*r := library.GetRouter(true)
	r.GET("/api/v1/atr/1", handler.AtrFindAll)

	req, _ := http.NewRequest("GET", "/api/v1/atr/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})*/

}
