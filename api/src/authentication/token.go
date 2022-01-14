package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken creates an authentication jwt token
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("this is a secret string, awesome!"))
}
