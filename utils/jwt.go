package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateTokenJWT(userPayload interface{}) (string, error) {
	//Creating Access Token
	expirationToken, _ := time.ParseDuration(os.Getenv("APP_TOKEN_EXPIRATION"))
	atClaim := jwt.MapClaims{
		"authorized": true,
		"user":       userPayload,
		"exp":        time.Now().Add(time.Minute * expirationToken).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)

	token, err := at.SignedString([]byte(os.Getenv("APP_TOKEN_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
