package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var sampleSecretKey = []byte("TokenSecretKey")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = username
	claims["exp"] = time.Now().Add(time.Minute * 6000).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
