package configs

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/atomic/atr/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	JwtActiveToken             *string
	PermissionID               *uint
	ActiveDB                   *gorm.DB
	BusinessID                 *string
	JwtActiveTokenMicroservice *string
)

type TempToken struct {
	JwtActiveTokenMicroservice string
	Data                       interface{}
}

// DBInit create connection to database
func DBInit() *models.DB {
	envname := "ATOMICGO"
	connection := os.Getenv(envname)
	fmt.Println(connection)
	fmt.Println("connection:" + connection)
	var connStr string
	connStr = connection + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		logrus.Error("failed to connect to database")
	}

	ActiveDB = db
	db.SingularTable(true)
	return &models.DB{DB: db}
}

// DBInitTest create connection to database
func DBInitTest() (*models.DB, sqlmock.Sqlmock, error, error) {
	var (
		dbtemp        *sql.DB
		DB            *gorm.DB
		mock          sqlmock.Sqlmock
		errmock       error
		errconnection error
	)
	dbtemp, mock, errmock = sqlmock.New()

	DB, errconnection = gorm.Open("mysql", dbtemp)

	DB.LogMode(true)

	db := models.DB{DB: DB}

	return &db, mock, errmock, errconnection
}
