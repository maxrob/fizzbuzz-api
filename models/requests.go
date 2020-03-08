package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstInt     int                `bson:"first_int" json:"first_int"`
	SecondInt    int                `bson:"second_int" json:"second_int"`
	Limit        int                `bson:"limit" json:"limit"`
	FirstString  string             `bson:"first_string" json:"first_string"`
	SecondString string             `bson:"second_string" json:"second_string"`
	Iteration    int                `bson:"iteration" json:"iteration"`
}
