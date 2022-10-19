package main

import (
	"log"
	"twittor/bd"
	"twittor/handlers"
)

func main() {

	if bd.ChequeoConection() == 0 {
		log.Fatalln("Sin conexion a bd")
		return
	}

	handlers.Manejadores()
}
