package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}
