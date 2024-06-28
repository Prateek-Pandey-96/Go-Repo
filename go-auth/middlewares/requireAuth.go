package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prateek69/go-auth/initializers"
	"github.com/prateek69/go-auth/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("Inside middleware")

	var val struct {
		token string
	}
	c.Bind(&val)
	print(val.token)
	token, err := jwt.Parse(val.token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var foundUser models.User
		initializers.DB.First(&foundUser, claims["sub"])

		if foundUser.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", foundUser)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}

//token - "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA2MjU5NjksInN1YmplY3QiOjJ9.PHRL0KLadXtGbUOdI1tAMfElZsJLAX6No-fr0ARaquI"
