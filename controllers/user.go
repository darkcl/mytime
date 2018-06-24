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

// UserController - User Logics
type UserController struct{}

func (u UserController) createJWT(user models.User) (*models.JwtToken, error) {
	// Issue Token
	config := config.GetConfig()
	secret := config.GetString("server.secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": string(user.ID),
	})
	tokenString, error := token.SignedString([]byte(secret))

	jwtToken := new(models.JwtToken)
	jwtToken.Token = tokenString

	return jwtToken, error
}

func (u UserController) getUserToken(userPayload forms.UserLoginForm) (*models.JwtToken, int, error) {
	var conn = db.GetDB()
	var user models.User
	if conn.Where("username = ?", userPayload.Username).First(&user).RecordNotFound() {
		return nil, http.StatusForbidden, errors.New("User and password does not match")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userPayload.Password)); err != nil {
		return nil, http.StatusForbidden, errors.New("User and password does not match")
	}
	token, err := u.createJWT(user)
	return token, http.StatusOK, err
}

func (u UserController) createUser(userPayload forms.UserSignupForm) (*models.JwtToken, int, error) {
	var conn = db.GetDB()
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	user := models.User{Username: userPayload.Username, Password: hash}

	if conn.Where("username = ?", userPayload.Username).First(&user).RecordNotFound() {
		if conn.NewRecord(user) {
			conn.Create(&user)

			token, err := u.createJWT(user)
			return token, http.StatusOK, err
		}
		return nil, http.StatusInternalServerError, errors.New("Cannot create user")
	}

	return nil, http.StatusForbidden, errors.New("User already exist")
}

// SignUp - Create a user
func (u UserController) SignUp(c *gin.Context) {
	var signUp forms.UserSignupForm
	jsonError := c.BindJSON(&signUp)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, jsonError)
		return
	}

	jwtToken, code, err := u.createUser(signUp)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(code, jwtToken)
	return
}

// GetToken - Get token
func (u UserController) GetToken(c *gin.Context) {
	var login forms.UserLoginForm
	jsonError := c.BindJSON(&login)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, jsonError)
		return
	}

	jwtToken, code, err := u.getUserToken(login)

	if err != nil {
		fmt.Println(err)
		c.JSON(code, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(code, jwtToken)
	return
}
