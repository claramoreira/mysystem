package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"system/cors"
)

const postsPath = "posts"

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		postList, err := getPostList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(postList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var post Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertPost(post)
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

func HandlePost(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", postsPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	postID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		post, err := getPost(postID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if post == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(post)
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
	postHandler := http.HandlerFunc(HandlePost)
	http.Handle(fmt.Sprintf("/%s/", postsPath), cors.Middleware(postHandler))
	postsHandler := http.HandlerFunc(HandlePosts)
	http.Handle(fmt.Sprintf("/%s", postsPath), cors.Middleware(postsHandler))
}
