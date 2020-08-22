package db

import (
	"context"
	"time"

	"github.com/gonzaloescobar/twitter-of-the-salty/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckUserAlreadyExists check user on dDB*/
func CheckUserAlreadyExists(mail string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	condition := bson.M{"mail": mail}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
