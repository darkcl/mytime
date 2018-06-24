package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/darkcl/mytime/config"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("authorization")
		config := config.GetConfig()
		secret := config.GetString("server.secret")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(secret), nil
				})
				if error != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, error)
					return
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					c.Set("userID", claims["id"])
					c.Next()
				} else {
					res := []string{"message", "Invalid Token"}
					c.AbortWithStatusJSON(http.StatusUnauthorized, res)
				}
			}
		} else {
			res := []string{"message", "An authorization header is required"}
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
		}
		c.Next()
	}
}
