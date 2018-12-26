package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var noteStore = make(map[string]Note)
var id int = 0

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{Addr: ":8080", Handler: r}
	log.Println("Listening to http://localhost:8080")
	server.ListenAndServe()
}
