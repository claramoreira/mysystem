package main

import (
	"log"
	"net/http"
	"system/database"
	"system/post"
	"system/topic"
	"system/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.SetupDatabase()
	post.SetupRoutes()
	user.SetupRoutes()
	topic.SetupRoutes()
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
