package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"system/cors"
)

var usersPath = "users"

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userList, err := getUsersList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(userList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", usersPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		user, err := getUser(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SetupRoutes() {
	usersHandler := http.HandlerFunc(HandleUsers)
	userHandler := http.HandlerFunc(HandleUser)
	http.Handle(fmt.Sprintf("/%s", usersPath), cors.Middleware(usersHandler))
	http.Handle(fmt.Sprintf("/%s/", usersPath), cors.Middleware(userHandler))
}
