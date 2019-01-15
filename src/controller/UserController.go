package controller

import (
	model "Structure/src/model"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var users []model.Users
	DB.Table("users").Scan(&users)
	json.NewEncoder(w).Encode(users)
}

func Store(w http.ResponseWriter, r *http.Request) {

	var users model.Users
	var data interface{}
	var request map[string]string
	json.NewDecoder(r.Body).Decode(&request)

	var response = make(map[string]interface{})

	staffId := request["staffId"]

	// Check user already exists
	DB.Table("users").Where("staff_id = ?", staffId).Find(&users)

	if users.StaffId != "" {
		response["status"] = "failed"
		response["msg"] = "Member already exists"
	} else {
		url := "http://103.206.184.11:90/auth_server/Project/Controllers/api-v1/staff_operation.php"
		payload := strings.NewReader("token=staff&staffid=" + staffId)
		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()

		json.NewDecoder(res.Body).Decode(&data)
		if data.(map[string]interface{})["isactive"].(bool) == true {
			record := data.(map[string]interface{})["record"].(map[string]interface{})

			users.Name = record["fullname"].(string)
			users.Email = record["email"].(string)
			users.StaffId = staffId
			users.Displayname = record["displayname"].(string)
			users.Desg = record["desg"].(string)
			users.Company = record["company"].(string)
			users.Dept = record["dept"].(string)
			users.Country = record["country"].(string)
			users.Location = record["location"].(string)
			users.Approved = "0"
			users.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			users.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

			// insert user to database
			DB.Create(&users)

			// add user to response to response
			response["status"] = "success"
			response["body"] = users
		} else {
			response["status"] = "failed"
			response["msg"] = "No member found"
		}
	}

	json.NewEncoder(w).Encode(response)
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

func TestUser(w http.ResponseWriter, r *http.Request) {

	var res map[string]interface{}

	json.NewDecoder(r.Body).Decode(&res)

	displayname := res["record"].(map[string]interface{})["displayname"].(string)
	w.Write([]byte(displayname))
	json.NewEncoder(w).Encode(res)

}
