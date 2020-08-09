package db

import (
	"context"
	"time"
	"github.com/gonzaloescobar/twitter-of-the-salty/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func InserUser(u models.User) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}