package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DevuelvoTweetsSeguidores struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID        string             `bson:"usuarioId" json:"usuarioId,omitempty"`
	UsuarioRelacioID string             `bson:"usuarioRelacionId" json:"usuarioRelacionId,omitempty"`
	Tweet            struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
