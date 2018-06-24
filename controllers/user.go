package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/darkcl/mytime/config"
	"github.com/darkcl/mytime/db"
	"github.com/darkcl/mytime/forms"
	"github.com/darkcl/mytime/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func createUser(userPayload forms.UserSignupForm) (*models.JwtToken, error) {
	var conn = db.GetDB()
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{Username: userPayload.Username, Password: hash}

	if conn.Where("username = ?", userPayload.Username).First(&user).RecordNotFound() {
		if conn.NewRecord(user) {
			// Issue Token
			conn.Create(&user)
			config := config.GetConfig()
			secret := config.GetString("server.secret")

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userName": userPayload.Username,
			})
			tokenString, error := token.SignedString([]byte(secret))

			jwtToken := new(models.JwtToken)
			jwtToken.Token = tokenString

			return jwtToken, error
		}
		return nil, errors.New("Cannot create user")
	}

	return nil, errors.New("User already exist")
}

func (u UserController) SignUp(c *gin.Context) {
	var signUp forms.UserSignupForm
	jsonError := c.BindJSON(&signUp)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, jsonError)
		return
	}

	jwtToken, err := createUser(signUp)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, jwtToken)
	return
}
