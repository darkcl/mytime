package server

import (
	"github.com/darkcl/mytime/controllers"
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
	}

	router.GET("/ping", health.Status)
	return router

}
