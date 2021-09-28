package topic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"system/cors"
)

const topicsPath = "topics"

func HandleTopics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		topicList, err := getTopicList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(topicList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var topic Topic
		err := json.NewDecoder(r.Body).Decode(&topic)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertTopic(topic)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleTopic(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", topicsPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	topicID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		topic, err := getTopic(topicID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if topic == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(topic)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodDelete:
		removeTopic(topicID)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupRoutes() {
	topicsHandler := http.HandlerFunc(HandleTopics)
	topicHandler := http.HandlerFunc(HandleTopic)
	http.Handle(fmt.Sprintf("/%s", topicsPath), cors.Middleware(topicsHandler))
	http.Handle(fmt.Sprintf("/%s/", topicsPath), cors.Middleware(topicHandler))
}
