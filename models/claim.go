package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim : This is the struct used to process the token*/
type Claim struct {
	Mail string             `json:"email"`
	ID   primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
