package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	k := mux.Vars(r)["id"]
	var noteToUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteToUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteToUpdate.CreatedOn = note.CreatedOn
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
	} else {
		log.Printf("Could not find any note with key %s to update", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	k := mux.Vars(r)["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	}else {
		log.Printf("Could not find any note with key %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}