package routers

import (
	"net/http"
	"twittor/bd"
	"twittor/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro id es obligatorio", 400)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio error al insertar relacion en BD"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
