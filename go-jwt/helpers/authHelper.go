package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/prateek69/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func IsAdmin(ctx *gin.Context, userId string) error {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("user_id")
	var err error
	if !((userType == "USER" && uid == userId) || (userType == "Admin")) {
		err = errors.New("unauthorized request! User doesnt have permission")
	}
	return err
}

func VerifyPassword(user *models.User, foundUser *models.User) error {
	err := bcrypt.CompareHashAndPassword([]byte((*user).Password), []byte((*foundUser).Password))
	if err != nil {
		return err
	}
	return nil
}
