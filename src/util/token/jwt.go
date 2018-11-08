package token

import (
	"config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetJWTToken(data map[string]interface{}) string {
	t := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := t.Claims.(jwt.MapClaims)
	for key, value := range data {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	jwtToken, _ := t.SignedString([]byte(config.Conf.Security.Secret))

	return jwtToken
}
