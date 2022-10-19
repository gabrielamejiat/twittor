package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"twittor/bd"
	"twittor/models"
)

var Email string
var IDUsuario string

func ProcesoToken(token string) (*models.Claim, bool, string, error) {

	miClave := []byte("MastersDelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("Token invalido")
	}

	return claims, false, "", err
}
