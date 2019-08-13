package middleware

import (
	"fmt"
	"net/http"

	"github.com/atomic/sip/src/helpers"

	"github.com/atomic/sip/configs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Auth function for checking JWT active session
func Auth(c *gin.Context) {
	var response interface{}
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	// if token.Valid && err == nil {
	if token != nil && err == nil {
		fmt.Println("token verified")

		configs.JwtActiveToken = &tokenString
	} else {
		response = helpers.Result{Status: "Warning", StatusCode: http.StatusUnauthorized, Message: "not authorized" + err.Error()}
		result := gin.H{
			"result": response,
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

}
