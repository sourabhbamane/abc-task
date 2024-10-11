package router

import (
	"healthclub/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() {
	r := mux.NewRouter()
	r.HandleFunc("/studio", controllers.Homepage).Methods("GET")
	r.HandleFunc("/studio/classes", controllers.CreateClass).Methods("POST")
	r.HandleFunc("/studio/bookings", controllers.BookClass).Methods("POST")

	//To listen to port
	log.Fatal(http.ListenAndServe(":8080", r))
}
