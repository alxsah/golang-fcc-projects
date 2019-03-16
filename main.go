package main

import (
  "log"
  "strconv"
	"net/http"
  "github.com/gorilla/mux"
  . "time"
  . "github.com/alxsah/fcc-timestamp-ms/utils"
)

func handleDate(w http.ResponseWriter, r *http.Request) {
  dateString := mux.Vars(r)["dateString"]
  res, err := handleDateString(dateString)
  if err != nil {
    res, err2 := handleUnixTimestamp(dateString)
    if err2 != nil {
      RespondWithError(w, http.StatusBadRequest, "Date could not be parsed")
      return
    }
    RespondWithJson(w, http.StatusOK, res)
    return
  }
  RespondWithJson(w, http.StatusOK, res)
}

func handleUnixTimestamp(ts string) (map[string]string, error) {
  t, err := strconv.ParseInt(ts, 10, 64)
  timeUnix := Unix(0, t * int64(Millisecond))
  timeUtc := timeUnix.UTC().Format(http.TimeFormat)
  return map[string]string{"unix": ts, "utc": timeUtc}, err
}

func handleDateString(ts string) (map[string]string, error) {
  layout := "2006-01-02"
  t, err := Parse(layout, ts)
  timeUnix := strconv.FormatInt(t.UnixNano() / int64(Millisecond), 10)
  timeUtc := t.UTC().Format(http.TimeFormat)
  return map[string]string{"unix": timeUnix, "utc": timeUtc}, err
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/api/timestamp/{dateString}", handleDate).Methods("GET")

  if err := http.ListenAndServe(":3000", r); err != nil {
    log.Fatal(err)
  }
}