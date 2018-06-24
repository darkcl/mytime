package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/darkcl/mytime/db"
	"github.com/darkcl/mytime/forms"
	"github.com/darkcl/mytime/models"
	"github.com/gin-gonic/gin"
)

// LeaveController - Leave Logic
type LeaveController struct{}

func (u LeaveController) createLeave(leavePayload forms.LeaveForm, userID uint64) (*models.Leave, int, error) {
	var conn = db.GetDB()
	layout := "01/02/2006"
	leaveDate, dateError := time.Parse(layout, leavePayload.Date)

	if dateError != nil {
		return nil, http.StatusBadRequest, errors.New("Date format invalid, example: 01/02/2006")
	}

	var user models.User
	if conn.Debug().Where("ID = ?", userID).First(&user).RecordNotFound() {
		return nil, http.StatusForbidden, errors.New("User not found")
	}

	leave := models.Leave{UserRefer: userID, LeaveDate: leaveDate, Reason: leavePayload.Reason}

	if conn.NewRecord(leave) {
		conn.Create(&leave)
		return &leave, http.StatusCreated, nil
	}

	return nil, http.StatusForbidden, errors.New("User already exist")
}

// CreateLeave - Create a user
func (u LeaveController) CreateLeave(c *gin.Context) {
	var leaveForm forms.LeaveForm
	userID, _ := c.Get("userID")

	jsonError := c.BindJSON(&leaveForm)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, jsonError)
		return
	}
	leave, code, err := u.createLeave(leaveForm, uint64(userID.(float64)))

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(code, leave)
	return
}
