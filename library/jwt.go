package library

import (
	"log"
	"time"

	"github.com/atomic/sip/configs"
	"github.com/atomic/sip/src/helpers"
	"github.com/dgrijalva/jwt-go"
)

const JwtSalt = "secret"

func JwtSignString(c helpers.Credential) (string, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)

	claims["ID"] = c.ID
	claims["Username"] = c.Username
	claims["Email"] = c.Email
	claims["Exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	} else {
		return token, nil
	}

}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "secret" // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func GetJWTClaims(token string) jwt.MapClaims {
	if token == "" {
		claims, _ := extractClaims(*configs.JwtActiveToken)
		return claims
	} else {
		claims, _ := extractClaims(token)
		return claims
	}
}

// func GetBusiness() string {
// 	claims := GetJWTClaims("")
// 	businessID := fmt.Sprintf("%v", claims["BusinessID"])
// 	strBusinessID := "business_id = " + businessID

// 	return strBusinessID
// }
