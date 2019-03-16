package main

import (
	. "github.com/alxsah/golang-fcc-projects/utils"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type FileMetadata struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	htmlString := `
	<!DOCTYPE html>
	<html>
		<body>
			<form action="api/fileanalyse" method="post" enctype="multipart/form-data">
			<input type="file" name="fileToUpload" id="fileToUpload">
			<input type="submit" value="Upload File" name="submit">
		</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, htmlString)
}

func handleFileAnalyse(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("fileToUpload")
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "File upload failed")
		return
	}
	defer file.Close()
	metadata := FileMetadata{
		header.Filename,
		header.Header.Get("Content-Type"),
		int(header.Size),
	}
	RespondWithJson(w, http.StatusOK, metadata)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/fileanalyse", handleFileAnalyse).Methods("POST")
	r.HandleFunc("/", handleRoot).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
