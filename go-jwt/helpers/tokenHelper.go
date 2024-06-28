package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/prateek69/go-jwt/database"
	"github.com/prateek69/go-jwt/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type TokenDetails struct {
	Email     string
	Name      string
	User_ID   string
	User_Type string
	jwt.StandardClaims
}

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateTokens(user *models.User) error {
	claims := TokenDetails{
		Email:     (*user).Email,
		Name:      (*user).Name,
		User_ID:   (*user).User_id,
		User_Type: (*user).User_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refresh_claims := TokenDetails{
		Email:     (*user).Email,
		Name:      (*user).Name,
		User_ID:   (*user).User_id,
		User_Type: (*user).User_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return err
	}
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return err
	}

	(*user).Token = token
	(*user).Refresh_token = refresh_token
	return nil
}

func HashPassword(user *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte((*user).Password), 14)
	if err != nil {
		return err
	}
	(*user).Password = string(bytes)
	return nil
}
