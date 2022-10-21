package routers

import (
	"io"
	"net/http"
	"os"
	"twittor/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro Id", 400)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usario no encontrado", 400)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)

	if err != nil {
		http.Error(w, "Imagen no encontrada", 400)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", 400)
		return
	}

}
