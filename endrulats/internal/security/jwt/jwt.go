package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(sessionname, sessionkey string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = sessionname
	claims["authorized"] = true
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(sessionkey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return t, nil
}
