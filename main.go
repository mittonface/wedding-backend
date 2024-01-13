package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/mittonface/wedding-backend/database"
	"github.com/mittonface/wedding-backend/middleware"
	"github.com/mittonface/wedding-backend/rsvp"
)
func handleRsvp(db database.RsvpDB, w http.ResponseWriter, r *http.Request){
	// decode the request body
	var rsvp rsvp.RSVP
	err := json.NewDecoder(r.Body).Decode(&rsvp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate the request body
	validate := validator.New()
	err = validate.Struct(rsvp)
	if err != nil {
		http.Error(w, "Invalid RSVP data", http.StatusBadRequest)
		return
	}

	// insert the rsvp
	err = db.InsertRSVP(&rsvp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	
}

func getRsvp(db database.RsvpDB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rsvpId := vars["rsvpId"]

	rsvp, err := db.GetRSVP(rsvpId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Println(rsvp)

	// Marshal the rsvp object into JSON
	jsonData, err := json.Marshal(rsvp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	w.Write(jsonData)
}

func health(db database.RsvpDB, w http.ResponseWriter, r *http.Request) {
	// Create a dummy RSVP
	dummyRsvp := rsvp.RSVP{
		RsvpId: "dummy",
		RsvpName: "dummy",
		NumGuests: 1,
		MealSelection1: "dummy",
		Attending: false,
	}

	// Insert the dummy RSVP into the database
	err := db.InsertRSVP(&dummyRsvp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a simple JSON structure
	response := map[string]string{"status": "OK"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	w.Write(jsonResponse)
}


func main() {

	db := &database.SupabaseDatabase{}
	err := db.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.Use(middleware.LogMiddleware) 

	r.HandleFunc("/rsvp", func(w http.ResponseWriter, r *http.Request) {
		handleRsvp(db, w, r)
	}).Methods("POST")
	r.HandleFunc("/rsvp/{rsvpId}", func(w http.ResponseWriter, r *http.Request) {
		getRsvp(db, w, r)
	}).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		health(db, w, r)
	}).Methods("GET")
	log.Println("Running server on :8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", r))
}
