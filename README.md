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

Set Env (Windows => User Variable)

  Variable => PORTATOMICGO || Value => :8080

  Variable => ATOMICGO || Value => root:@(localhost:3306)/atr 

  Variable => BASICTOKENGO || Value => {"Type":"mobileapp","SecretKey":"Angke Fren","Key":"LKHlhb899Y09olUi"}]

  Variable => WEBHOOKSLACK || Value => https://hooks.slack.com/services/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


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