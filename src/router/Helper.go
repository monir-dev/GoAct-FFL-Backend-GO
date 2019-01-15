package router

import (
	"Structure/src/config/db"
	jwt "Structure/src/middleware"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var DB = db.Connect()
var USER_ID int

func IsAuthenticated(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearer := r.Header.Get("Authorization")
		if bearer != "" {
			token := strings.Replace(bearer, "Bearer ", "", 1)
			response, err := jwt.PurseToken(token)

			checkErr(err)

			// set user id
			user_id, err := strconv.Atoi(response)
			checkErr(err)
			USER_ID = user_id

			// process request
			endpoint(w, r)

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		log.Println("Request URI : ", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
