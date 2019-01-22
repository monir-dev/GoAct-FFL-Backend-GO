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
	r.Handle("/auth-user-bulk-delete", IsAuthenticated(controller.DeleteBuldAuthUser)).Methods("POST")
	r.Handle("/auth-users-assign-role/{id}", IsAuthenticated(controller.AssignAuthUserRole)).Methods("POST")


	r.Handle("/roles", IsAuthenticated(controller.GetRoles)).Methods("GET")
	r.Handle("/roles", IsAuthenticated(controller.AddRole)).Methods("POST")
	r.Handle("/roles-bulk-delete", IsAuthenticated(controller.DeleteBulkRoles)).Methods("POST")
	r.Handle("/roles/{id}", IsAuthenticated(controller.EditRole)).Methods("PUT")
	r.Handle("/roles/{id}", IsAuthenticated(controller.DeleteRole)).Methods("DELETE")

	r.Handle("/permissions/{role_id}", IsAuthenticated(controller.GetPermissions)).Methods("GET")
	r.Handle("/permissions/{role_id}", IsAuthenticated(controller.AssignPermissionsToRole)).Methods("POST")

	r.HandleFunc("/test", controller.Testing).Methods("GET")
	r.Handle("/check", IsAuthenticated(CheckHandler))
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(USER_ID)))
}
