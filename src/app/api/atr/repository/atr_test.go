package repository_test

import (
	"testing"

	"github.com/atomic/atr/configs"
	"github.com/atomic/atr/models"
	"github.com/atomic/atr/src/app/api/atr/repository"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type AtrSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository repository.AtrRepository
	atr        *models.Atr

	VariableThatShouldStartAtFive int
}

func (s *AtrSuite) SetupTest() {

	db1, mock, errmock, errconnection := configs.DBInitTest()
	require.NoError(s.T(), errmock)
	require.NoError(s.T(), errconnection)

	s.mock = mock
	s.DB = db1.DB

	s.repository = repository.NewAtrRepository(db1)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRunAll(t *testing.T) {
	suite.Run(t, new(AtrSuite))
}
