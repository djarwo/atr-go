# README ATR-BPNP#

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

CREATE ENVIRONMENT
  Environment Name: ATR
  Windows: 
    ATR = path\to\your\environment\location
  LINUX
    export ATR=path\to\your\environment\location

Configuration

  go get -u "github.com/gin-gonic/gin"

  go get -u "github.com/jinzhu/gorm"

  go get -u "github.com/jinzhu/gorm/dialects/mysql"

  go get -u "github.com/go-sql-driver/mysql"

  go get -u "github.com/NaySoftware/go-fcm"

  go get -u "github.com/maddevsio/fcm"

  go get -u "github.com/dgrijalva/jwt-go"

  go get -u "github.com/gin-contrib/cors"

  go get -u "github.com/spf13/viper"

  go get -u "github.com/sirupsen/logrus"

  go get gopkg.in/DATA-DOG/go-sqlmock.v1

Unit Testing
  go get -u "github.com/vektra/mockery/.../", then $GOPATH/bin/mockery

// Postman Collection 
https://www.getpostman.com/collections/b81ec19a76f6f5428a7f


Install to Bin

go install path/to/your/code/project_name/main.go

Run 

.\bin\main 


Run program without installation

cd to path/to/your/code/project_name/

then run main.go

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact