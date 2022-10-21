package bd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"twittor/models"
)

func ConsultoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioId":         t.UsuarioID,
		"usuarioRelacionId": t.UsuarioRelacionID,
	}

	var resultados models.Relacion
	fmt.Println(resultados)

	err := col.FindOne(ctx, condicion).Decode(&resultados)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil

}
