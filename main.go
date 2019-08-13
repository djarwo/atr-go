package main

import (
	"fmt"
	"os"

	"github.com/atomic/atr/src/helpers"

	"github.com/atomic/atr/configs"
	"github.com/atomic/atr/src/routes"

	"github.com/spf13/viper"
)

//Init function for initialize config
func init() {
	envname := "ATR-BPN"
	path := os.Getenv(envname)
	fmt.Println(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		panic(envname + " PATH NOT FOUND (" + path + ")")
	}

	viper.SetConfigFile(path + `config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

//Main function for start entry golang
func main() {

	db := configs.DBInit()

	errMigrate := helpers.DBMigrate(db)
	if errMigrate != nil {
		fmt.Println("Migrate Error")
	}

	routes.RegisterRoutes(db)
}
