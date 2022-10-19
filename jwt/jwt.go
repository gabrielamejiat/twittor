package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"twittor/models"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersDelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokeStr, err := token.SignedString(miClave)

	if err != nil {
		return tokeStr, err
	}

	return tokeStr, nil

}
