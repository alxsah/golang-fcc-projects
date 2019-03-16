package main

import (
	"encoding/json"
	. "github.com/alxsah/golang-fcc-projects/utils"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"strconv"
)

type ShortURLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    int    `json:"short_url"`
}

var shortUrlMap = make(map[int]string)
var mapCounter int

func handleNewUrl(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	_, err := net.LookupIP(body["url"])
	if err != nil {
		RespondWithJson(w, http.StatusBadRequest, map[string]string{"error": "invalid URL"})
		return
	}
	shortUrlMap[mapCounter] = body["url"]
	RespondWithJson(w, http.StatusCreated, ShortURLResponse{body["url"], mapCounter})
	mapCounter++
}

func handleShortUrl(w http.ResponseWriter, r *http.Request) {
	intId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Could not parse ID")
		return
	}
	redirectUrl := shortUrlMap[intId]
	if redirectUrl != "" {
		http.Redirect(w, r, "https://"+redirectUrl, http.StatusSeeOther)
	}
}

func main() {
	mapCounter = 0
	r := mux.NewRouter()
	r.HandleFunc("/api/shorturl/new", handleNewUrl).Methods("POST")
	r.HandleFunc("/api/shorturl/{id}", handleShortUrl).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
