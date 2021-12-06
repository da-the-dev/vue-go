package main

import (
	"backend/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("Hi")
	}).Methods("GET")

	router.HandleFunc("/posts", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(db.DBRead())
	}).Methods("GET")

	println("Server started at port :3000. Connect: http://localhost:3000")
	http.ListenAndServe(":3000", router)
}
