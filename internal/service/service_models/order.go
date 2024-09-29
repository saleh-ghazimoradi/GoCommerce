package service_models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	OrderID       primitive.ObjectID `bson:"_id"`
	OrderCart     []ProductUser      `json:"order_list" bson:"order_list"`
	OrderAt       time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price         int                `json:"total_price" bson:"total_price"`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}
