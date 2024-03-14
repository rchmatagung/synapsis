package utils

import (
	"synapsis/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenereateJWT(conf *config.Config, customerId int64, employeeName string) (string, error) {
	atSecretKey := conf.Authorization.JWT.AccessTokenSecretKey

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["customer_id"] = customerId
	claims["username"] = employeeName
	claims["exp"] = time.Now().Add(time.Hour * 24 * conf.Authorization.JWT.AccessTokenDuration).Unix()
	tokenString, err := token.SignedString([]byte(atSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
