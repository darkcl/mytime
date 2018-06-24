package models

import (
	"github.com/darkcl/mytime/config"
	"github.com/darkcl/mytime/forms"
	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

func (h User) Signup(userPayload forms.UserSignupForm) (*JwtToken, error) {

	config := config.GetConfig()
	secret := config.GetString("server.secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": userPayload.Username,
		"password": userPayload.Password,
	})
	tokenString, error := token.SignedString([]byte(secret))

	jwtToken := new(JwtToken)
	jwtToken.Token = tokenString

	return jwtToken, error
}
