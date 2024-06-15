package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id     primitive.ObjectID `bson:"_id"`
	UserId int                `bson:"userid" json:"userid"`
	Name   string             `json:"name"`
	Gender string             `json:"gender"`
	Age    int                `json:"age"`
}
