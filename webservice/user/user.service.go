package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func SetupRoutes() {
	usersHandler := http.HandlerFunc(HandleUsers)
	http.Handle(fmt.Sprintf("/%s", usersPath), usersHandler)
}
