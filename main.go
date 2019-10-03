package main

import (
	"os"
	"time"

	"github.com/atomic/atr/configs"
	"github.com/atomic/atr/src/routes"
)

var loc *time.Location

//Init function for initialize config
func init() {

}

//Main function for start entry golang
func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	db := configs.DBInit()

	routes.RegisterRoutes(db)

}
