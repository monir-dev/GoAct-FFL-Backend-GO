package router

import (
	controller "Structure/src/Controller"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", controller.Register).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	r.Handle("/check", IsAuthenticated(CheckHandler))
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(USER_ID)))
}
