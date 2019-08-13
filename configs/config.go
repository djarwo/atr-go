package configs

import (
	"database/sql"
	"fmt"

	"github.com/atomic/atr/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	JwtActiveToken *string
	ActiveDB       *gorm.DB
)

// DBInit create connection to database
func DBInit() *models.DB {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	var connStr string
	connStr = "" + dbUser + ":" + dbPass + "@(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println(connStr)

	db, err := gorm.Open("mysql", connStr)
	//gorm.Open("mysql", "dev:seven@(localhost:3306)/sip?charset=utf8&parseTime=True&loc=Local")
	//gorm.Open("mysql", "dev:@seven(127.0.0.1:3306)/sip?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	// fmt.Println("Masuk InitDB")
	ActiveDB = db

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
