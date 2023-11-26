package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `bson:"user_name,omitempty"`
	Password string             `bson:"password,omitempty"`
	Contact  Contact            `bson:"contact,omitempty"`
}

type Contact struct {
	Email     string `bson:"email,omitempty"`
	TelNo     string `bosn:"telno,omitempty"`
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
}
