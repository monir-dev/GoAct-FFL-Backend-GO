package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func WebRoutes(r *mux.Router) {
	r.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
