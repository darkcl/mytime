package server

import (
	"github.com/darkcl/mytime/controllers"
	"github.com/darkcl/mytime/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.POST("", user.SignUp)
			userGroup.POST("/token", user.GetToken)
		}

		meGroup := v1.Group("me")
		{
			leaveCtrl := new(controllers.LeaveController)
			meGroup.Use(middlewares.AuthMiddleware())
			meGroup.GET("/leaves", leaveCtrl.ListLeave)
			meGroup.POST("/leaves", leaveCtrl.CreateLeave)
			meGroup.DELETE("/leaves/:leaveId", leaveCtrl.DeleteLeave)
		}
	}

	router.GET("/ping", health.Status)
	return router

}
