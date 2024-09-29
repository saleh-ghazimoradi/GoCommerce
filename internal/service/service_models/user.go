package service_models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"first_name" validate:"required,min=2,max=30"`
	LastName       *string            `json:"last_name"  validate:"required,min=2,max=30"`
	Password       *string            `json:"password"   validate:"required,min=6"`
	Email          *string            `json:"email"      validate:"email,required"`
	Phone          *string            `json:"phone"      validate:"required"`
	Token          *string            `json:"token"`
	RefreshToken   *string            `json:"refresh_token"`
	CreateAt       time.Time          `json:"create_at"`
	UpdatedAt      time.Time          `json:"update_at"`
	UserID         string             `json:"user_id"`
	UserCart       []ProductUser      `json:"usercart" bson:"usercart"`
	AddressDetails []Address          `json:"address" bson:"address"`
	OrderStatus    []Order            `json:"orders" bson:"orders"`
}
