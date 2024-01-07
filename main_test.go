package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mittonface/wedding-backend/database"
	"github.com/mittonface/wedding-backend/rsvp"
)


func TestHandleRsvp(t *testing.T) {
	// Mock the SupabaseDatabase
	mockDB := &database.MockSupabaseDatabase{}

	// Prepare the request
	rsvp := rsvp.RSVP{
		RsvpId: "test",
		RsvpName: "test",
		NumGuests: 1,
		MealSelection1: "test",
		Attending: true,
	}
	rsvpBytes, _ := json.Marshal(rsvp)
	req, err := http.NewRequest("POST", "/rsvp", bytes.NewBuffer(rsvpBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handleRsvp(mockDB, rr, req)

	// Assert the status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Assert that the RSVP was inserted correctly
	if mockDB.Rsvp == nil || mockDB.Rsvp.RsvpId != rsvp.RsvpId {
		t.Errorf("RSVP was not inserted correctly")
	}
}


func TestHealth(t *testing.T) {
	// Mock the SupabaseDatabase
	mockDB := &database.MockSupabaseDatabase{}

	// Prepare the request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Call the handler
	health(mockDB, rr, req)

	// Assert the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}