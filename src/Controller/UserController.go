package controller

import (
	model "Structure/src/Model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var users []model.Users
	DB.Table("users").Scan(&users)
	json.NewEncoder(w).Encode(users)
}

func ChangeUserApproveStatus(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)

	DB.Table("users").Where("id = ?", id).Update("approved", data["approved"])

	var response = make(map[string]string, 0)
	response["status"] = "success"
	json.NewEncoder(w).Encode(response)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	var users []model.Users

	DB.Table("users").Where("id = ?", id).Delete(&users)

	var response = make(map[string]string, 0)
	response["status"] = "success"
	json.NewEncoder(w).Encode(response)
}

func UserPermissions(w http.ResponseWriter, r *http.Request) {
	var results []map[string]interface{}
	var uid, rid int
	var uname, email, rname string

	rows, err := DB.Table("users").Select("users.id, users.name, users.email, roles.id, roles.name").Joins("left join role_user on role_user.user_id = users.id").Joins("left join roles on roles.id = role_user.role_id").Rows()
	checkErr(err)

	for rows.Next() {
		rows.Scan(&uid, &uname, &email, &rid, &rname)

		result := map[string]interface{}{
			"user_id":    uid,
			"user_name":  uname,
			"user_email": email,
			"role_id":    rid,
			"role_name":  rname,
		}

		results = append(results, result)
	}

	json.NewEncoder(w).Encode(results)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
