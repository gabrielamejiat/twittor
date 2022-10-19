package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoCN = ConectarBd()
var clientOptions = options.Client().ApplyURI("mongodb+srv://marmejia:1020846Mar$@twittor.edjkvud.mongodb.net/test")

// ConectarBd es la función que permite conectar a la bd
func ConectarBd() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalln(err.Error())
		return client

	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalln(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")

	return client
}

//ChequeoConection es pong a bd
func ChequeoConection() int {

	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}
	return 1
}
