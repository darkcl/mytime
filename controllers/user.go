package controllers

import (
	"net/http"

	"github.com/darkcl/mytime/forms"
	"github.com/darkcl/mytime/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) SignUp(c *gin.Context) {
	var signUp forms.UserSignupForm
	jsonError := c.BindJSON(&signUp)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, jsonError)
		return
	}

	jwtToken, err := userModel.Signup(signUp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to signup", "error": err})
		c.Abort()
		return
	}
	c.JSON(200, jwtToken)
	return
}
