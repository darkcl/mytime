package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model

	Username string `json:"userName"`
	Password []byte `gorm:"type:varchar(255); not null"`
}

type JwtToken struct {
	Token string `json:"token"`
}

// func (h User) Signup(userPayload forms.UserSignupForm) (*JwtToken, error) {

// 	config := config.GetConfig()
// 	secret := config.GetString("server.secret")

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"userName": userPayload.Username,
// 		"password": userPayload.Password,
// 	})
// 	tokenString, error := token.SignedString([]byte(secret))

// 	jwtToken := new(JwtToken)
// 	jwtToken.Token = tokenString

// 	return jwtToken, error
// }
