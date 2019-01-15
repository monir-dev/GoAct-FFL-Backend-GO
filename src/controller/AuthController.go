package controller

import (
	"Structure/src/config"
	jwt "Structure/src/middleware"
	"Structure/src/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user model.User
	var response = make(map[string]interface{})

	json.NewDecoder(r.Body).Decode(&user)

	// For Form data
	// var email = r.FormValue("email")
	// var password = r.FormValue("password")

	password := user.Password

	DB.Where("email = ?", user.Email).First(&user)

	if len(password) == len(user.Password) {

		response["status"] = "failed"
		response["msg"] = "No user is associated with this email"
		json.NewEncoder(w).Encode(response)
		return

	} else {

		// check password
		var hash config.Hash
		if err := hash.Compare(user.Password, password); err != nil {

			response["status"] = "failed"
			response["msg"] = "Invalid Username and Password"
			json.NewEncoder(w).Encode(response)
			return
		}

		// TODO - also need to add role and permissions
		userInfo, err := jwt.CreateJwtToken(user.ID, user.Name, user.Email)
		checkErr(err)

		json.NewEncoder(w).Encode(userInfo)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

	var user model.User

	// decode resquest body
	json.NewDecoder(r.Body).Decode(&user)

	// validations
	if user.Name == "" {
		SendErrorResponse(w, map[string]interface{}{
			"name": "Name is required",
		}, 404)
		return
	}

	if user.Email == "" {
		SendErrorResponse(w, map[string]interface{}{
			"email": "email is required",
		}, 404)
		return
	}

	// check if user already exist
	count := 0
	DB.Where("email = ?", user.Email).Find(&user).Count(&count)
	if count > 0 {
		SendErrorResponse(w, map[string]interface{}{
			"email": "Email Already Exists",
		}, 404)
		return
	}

	// if user doesn't exits hass password then create user

	// hash password
	var hash config.Hash
	hashedPassword, err := hash.Generate(user.Password)
	checkErr(err)

	// add hashed password to user
	user.Password = hashedPassword

	// create user
	DB.Create(&user)

	// return response
	SendSuccessResponse(w, user)
	return
	// user.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	// user.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	DB.Table("users").Scan(&users)
	json.NewEncoder(w).Encode(users)
}

func UpdateAuthUser(w http.ResponseWriter, r *http.Request) {

	var user model.User
	var checkUser model.User

	// get query id
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	checkErr(err)

	// decode resquest body
	json.NewDecoder(r.Body).Decode(&user)

	// set user id
	user.ID = id

	// validations
	if user.Name == "" {
		SendErrorResponse(w, map[string]interface{}{
			"name": "Name is required",
		}, 404)
		return
	}

	if user.Email == "" {
		SendErrorResponse(w, map[string]interface{}{
			"email": "email is required",
		}, 404)
		return
	}

	// check if user already exist
	count := 0
	DB.Where("id != ? AND email = ?", id, user.Email).Find(&checkUser).Count(&count)
	if count > 0 {
		SendErrorResponse(w, map[string]interface{}{
			"email": "Email Already Exists",
		}, 404)
		return
	}

	// if user doesn't exits hass password then create user

	if user.Password != "" {
		// hash password
		var hash config.Hash
		hashedPassword, err := hash.Generate(user.Password)
		checkErr(err)

		// add hashed password to user
		user.Password = hashedPassword
	}



	// Update user
	DB.Save(&user)

	// return response
	SendSuccessResponse(w, user)
	return
}

func DeleteAuthUser(w http.ResponseWriter, r *http.Request)  {
	id := mux.Vars(r)["id"]
	var user []model.User

	DB.Table("users").Where("id = ?", id).Delete(&user)

	SendSuccessResponse(w, make(map[string]string, 0))
}