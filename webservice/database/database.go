package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DBConfig struct {
	Host     string `json:"host"`
	Database string `json:"db"`
	User     string `json:"user"`
	Password string `json:"pass"`
}

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	var db DBConfig
	jsonFile, err := os.Open("./database/.info.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &db)
	if err != nil {
		log.Fatal(err)
	}
	DbConn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", db.User, db.Password, db.Host, db.Database))
	if err != nil {
		log.Fatal(err)
	}
	// DbConn.SetMaxOpenConns(4)
	// DbConn.SetMaxIdleConns(4)
	// DbConn.SetConnMaxLifetime(60 * time.Second)
}
