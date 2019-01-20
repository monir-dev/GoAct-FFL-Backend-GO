package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {

	// Middleware applicable for all request
	r.Use(CommonMiddleware)

	// Structure/src/router/AuthRoutes.go
	// Auth, Roles, Permission
	AuthRoutes(r)

	// Structure/src/router/UsersRoutes.go
	UsersRoutes(r)

	// Structure/src/router/WebRoutes.go
	WebRoutes(r)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
