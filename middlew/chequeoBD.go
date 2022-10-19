package middlew

import (
	"net/http"
	"twittor/bd"
)

//Los handlers reciben y devuelven lo mismo
func ChequeoBd(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConection() == 0 {
			http.Error(writer, "Conexion perdida con la bd", 500)
			return
		}

		next.ServeHTTP(writer, r)
	}
}
