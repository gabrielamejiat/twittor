package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"twittor/models"
)

func InsertoTweet(t models.GraboTweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userId":  t.UserId,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)

	return ObjId.String(), true, nil

}
