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
		http.Error(w, fmt.Sprintf("Invalid RSVP Data: %s", err), http.StatusBadRequest)
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

func getAllRsvps(db database.RsvpDB, w http.ResponseWriter, r *http.Request) {

	rsvps, err := db.GetAllRSVPs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Marshal the rsvp object into JSON
	jsonData, err := json.Marshal(rsvps)
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
		Added: "2024-01-07 04:30:29.905495",
	}

	// Insert the dummy RSVP into the database
	err := db.InsertRSVP(&dummyRsvp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.DeleteRSVPs(dummyRsvp.RsvpId)
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
	r.Use(middleware.EnableCors) 

	r.PathPrefix("/").Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		origin := r.Header.Get("Origin")
		if origin == "http://localhost:3000" || origin == "https://jessandbrent.ca" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// Handle OPTIONS request here
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})
	r.HandleFunc("/rsvps", func(w http.ResponseWriter, r *http.Request) {
		getAllRsvps(db, w, r)
	}).Methods("GET")
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
	// log.Fatal(http.ListenAndServe(":8080", r))
}
