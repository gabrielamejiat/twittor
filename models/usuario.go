package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatars" json:"avatars,omitempty"`
	Banner          string             `bson:"banners" json:"banners,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
}
