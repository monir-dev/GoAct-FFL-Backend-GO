package controller

import (
	model "Structure/src/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {

	var roles []map[string]interface{}

	var id, created_by int
	var name, display_name, description string

	rows, err := DB.Table("roles").Select("id, name, display_name, description, created_by").Rows()
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
		rows.Scan(&id, &name, &display_name, &description, &created_by)

		role := map[string]interface{}{
			"id":           id,
			"name":         name,
			"display_name": display_name,
			"description":  description,
			"created_by":   created_by,
		}

		roles = append(roles, role)
	}
	json.NewEncoder(w).Encode(roles)
}

func AddRole(w http.ResponseWriter, r *http.Request) {
	var role model.Role
	json.NewDecoder(r.Body).Decode(&role)

	DB.Create(&role)

	var response = make(map[string]interface{})
	// add role to response to response
	response["status"] = "success"
	response["body"] = role

	json.NewEncoder(w).Encode(response)
}

func EditRole(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)

	DB.Table("roles").Where("id = ?", id).Update(data)

	var response = make(map[string]string, 0)
	response["status"] = "success"
	json.NewEncoder(w).Encode(response)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	DB.Exec("DELETE FROM roles where id = ?", id)

	var response = make(map[string]string, 0)
	response["status"] = "success"
	json.NewEncoder(w).Encode(response)
}