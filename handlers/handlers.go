package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twittor/middlew"
	"twittor/routers"
)

func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBd(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBd(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlew.ChequeoBd(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBd(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBd(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBd(middlew.ValidoJWT(routers.LeoTweet))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatalln(http.ListenAndServe(":"+PORT, handler))

}
