package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type User struct {
	ID			primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	FirstName	string 			   `bson:"firstName" json:"firstName,omitempty"`
	LastName	string 			   `bson:"lastName" json:"lastName,omitempty"`
	BirthDate	time.Time 		   `bson:"birthDate" json:"birthDate,omitempty"`
	Mail		string 			   `bson:"mail" json:"mail,omitempty"`
	Password 	string			   `bson:"password" json:"password,omitempty"`
	Avatar		string			   `bson:"avatar" json:"avatar,omitempty"`
	Banner		string			   `bson:"banner" json:"banner,omitempty"`
	Biography	string			   `bson:"biography" json:"biography,omitempty"`
	Location 	string			   `bson:"location" json:"location,omitempty"`
	Website		string			   `bson:"website" json:"website,omitempty"`

}