package controller

import (
	"Structure/src/config"
	"Structure/src/config/db"
	"encoding/json"
	"net/http"
)

var DB = db.Connect()

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SendErrorResponse(w http.ResponseWriter, data map[string]interface{}, code int) {
	var res = make(map[string]interface{})
	w.WriteHeader(code)
	res = data
	res["status"] = "failed"
	res["code"] = code
	json.NewEncoder(w).Encode(res)
}

func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	var res = make(map[string]interface{})
	w.WriteHeader(200)
	res["status"] = "success"
	res["code"] = 200
	res["data"] = data
	json.NewEncoder(w).Encode(res)
}

func Testing(w http.ResponseWriter, r *http.Request) {
	var hash config.Hash

	hashedPass := "$2a$10$0r2e2fqgMF8D8Rbrs8DcKOqUzsuaETSNh95nRScEbKTMyn4ctO8Ia"

	password := "monir"
	// hashPass, err := hash.Generate(password)
	// checkErr(err)

	if err := hash.Compare(hashedPass, password); err != nil {
		w.Write([]byte("not Matched"))
	}

	w.Write([]byte("Matched"))

}
