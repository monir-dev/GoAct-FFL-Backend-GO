package router

import (
	controller "Structure/src/controller"

	"github.com/gorilla/mux"
)

func UsersRoutes(r *mux.Router) {
	r.Handle("/users", IsAuthenticated(controller.Index)).Methods("GET")
	r.Handle("/users", IsAuthenticated(controller.Store)).Methods("POST")
	r.Handle("/users/{id}", IsAuthenticated(controller.ChangeUserApproveStatus)).Methods("PUT")
	r.Handle("/users/{id}", IsAuthenticated(controller.Delete)).Methods("DELETE")
	r.Handle("/permissions", IsAuthenticated(controller.UserPermissions)).Methods("GET")

	r.Handle("/users/test", IsAuthenticated(controller.TestUser)).Methods("GET")

	r.Handle("/roles", IsAuthenticated(controller.GetRoles)).Methods("GET")
	r.Handle("/roles", IsAuthenticated(controller.AddRole)).Methods("POST")
	r.Handle("/roles/{id}", IsAuthenticated(controller.EditRole)).Methods("PUT")
	r.Handle("/roles/{id}", IsAuthenticated(controller.DeleteRole)).Methods("DELETE")
}
