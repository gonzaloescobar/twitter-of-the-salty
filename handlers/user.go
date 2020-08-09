package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gonzaloescobar/twitter-of-the-salty/middlew"
	"github.com/gonzaloescobar/twitter-of-the-salty/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto el Handler y pongo a escuhar al Servidor*/
func Handlers(){
	router := mux.NewRouter()

	router.HandleFunc("/enrollment", middlew.CheckDB(routers.User)).Methods("POST")
	 
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}