package router

import (
	controller "Structure/src/controller"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	r.Handle("/auth-users", IsAuthenticated(controller.GetAllUser)).Methods("GET")
	r.Handle("/auth-user/{id}", IsAuthenticated(controller.UpdateAuthUser)).Methods("POST")
	r.Handle("/auth-user/{id}", IsAuthenticated(controller.DeleteAuthUser)).Methods("DELETE")


	r.HandleFunc("/test", controller.Testing).Methods("GET")
	r.Handle("/check", IsAuthenticated(CheckHandler))
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(USER_ID)))
}
