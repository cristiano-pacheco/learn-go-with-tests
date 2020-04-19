package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Oh Yeah!\n"))
}

func store(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(body, &newEvent)
	events = append(events, newEvent)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func show(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedEvent event

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(body, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == id {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i, singleEvent := range events {
		if singleEvent.ID == id {
			events = append(events[:i], events[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/events", store).Methods("POST")
	router.HandleFunc("/events", index).Methods("GET")
	router.HandleFunc("/events/{id}", show).Methods("GET")
	router.HandleFunc("/events/{id}", update).Methods("PUT")
	router.HandleFunc("/events/{id}", destroy).Methods("DELETE")

	router.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
