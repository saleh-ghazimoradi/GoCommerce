package service_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Rating      *uint              `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}
