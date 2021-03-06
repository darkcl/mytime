package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/darkcl/mytime/db"
	"github.com/darkcl/mytime/forms"
	"github.com/darkcl/mytime/models"
	"github.com/gin-gonic/gin"
)

// LeaveController - Leave Logic
type LeaveController struct{}

func (u LeaveController) createLeave(leavePayload forms.LeaveForm, userID string) (*models.Leave, int, error) {
	var conn = db.GetDB()
	layout := "01/02/2006"
	leaveDate, dateError := time.Parse(layout, leavePayload.Date)

	if dateError != nil {
		return nil, http.StatusBadRequest, errors.New("Date format invalid, example: 01/02/2006")
	}

	var user models.User
	fmt.Println(userID)
	ID, _ := strconv.ParseUint(userID, 10, 64)

	if conn.Debug().Where("ID = ?", ID).First(&user).RecordNotFound() {
		return nil, http.StatusForbidden, errors.New("User not found")
	}

	leave := models.Leave{LeaveDate: leaveDate}
	if conn.First(&leave).RecordNotFound() == false {
		return nil, http.StatusForbidden, fmt.Errorf("Already leave on this day, Leave ID = %d", leave.ID)
	}

	result := models.Leave{UserRefer: ID, LeaveDate: leaveDate, Reason: leavePayload.Reason}

	if conn.NewRecord(result) {
		conn.Create(&result)
		return &result, http.StatusCreated, nil
	}

	return nil, http.StatusForbidden, errors.New("User already exist")
}

func (u LeaveController) deleteLeave(leaveID string) (int, error) {
	var conn = db.GetDB()
	var leave models.Leave
	if conn.Where("ID = ?", leaveID).First(&leave).RecordNotFound() {
		return http.StatusForbidden, errors.New("Leave day not found")
	}

	conn.Delete(leave)
	return http.StatusNoContent, nil
}

// CreateLeave - Add a day off
func (u LeaveController) CreateLeave(c *gin.Context) {
	var leaveForm forms.LeaveForm
	ID, _ := c.Get("userID")
	userID, _ := ID.(string)

	jsonError := c.BindJSON(&leaveForm)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, jsonError)
		return
	}
	leave, code, err := u.createLeave(leaveForm, userID)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	leaveResponse := models.LeaveResponse{LeaveID: fmt.Sprint(leave.ID), LeaveDate: leave.LeaveDate, Reason: leave.Reason}
	c.JSON(code, leaveResponse)
	return
}

// ListLeave - List All Leaves
func (u LeaveController) ListLeave(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
	return
}

// DeleteLeave - Reave a day off
func (u LeaveController) DeleteLeave(c *gin.Context) {
	leaveID := c.Param("leaveId")
	code, err := u.deleteLeave(leaveID)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Status(code)
	return
}
