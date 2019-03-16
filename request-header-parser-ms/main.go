package main

import (
	. "github.com/alxsah/golang-fcc-projects/utils"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type RequestData struct {
	IPAddress string `json:"ipaddress"`
	Language  string `json:"language"`
	Software  string `json:"software"`
}

func parseHeaders(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Server error")
	}
	language := r.Header.Get("Accept-Language")
	sysinfo := r.Header.Get("User-Agent")
	rdo := RequestData{host, language, sysinfo}
	RespondWithJson(w, http.StatusOK, rdo)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/whoami", parseHeaders).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
