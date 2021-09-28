package community

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"system/cors"
)

var communitiesPath = "communities"

func HandleCommunities(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		communityList, err := getCommunitiesList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(communityList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var community Community
		err := json.NewDecoder(r.Body).Decode(&community)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertCommunity(community)
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

func HandleCommunity(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", communitiesPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	communityID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		community, err := getCommunity(communityID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if community == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(community)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPut:
		var community Community
		err := json.NewDecoder(r.Body).Decode(&community)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if community.CommunityID != communityID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateCommunity(community)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		removeCommunity(communityID)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupRoutes() {
	communitiesHandler := http.HandlerFunc(HandleCommunities)
	communityHandler := http.HandlerFunc(HandleCommunity)
	http.Handle(fmt.Sprintf("/%s", communitiesPath), cors.Middleware(communitiesHandler))
	http.Handle(fmt.Sprintf("/%s/", communitiesPath), cors.Middleware(communityHandler))
}
