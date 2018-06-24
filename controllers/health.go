package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
	Message string `json:"message"`
}

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	msg := new(Health)
	msg.Message = "OK"
	c.JSON(http.StatusOK, msg)
}
