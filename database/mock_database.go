package database

import (
	"errors"

	"github.com/mittonface/wedding-backend/rsvp"
)

type MockSupabaseDatabase struct {
	Rsvp *rsvp.RSVP
}

func (db *MockSupabaseDatabase) Initialize() error {
	return nil
}

func (db *MockSupabaseDatabase) InsertRSVP(rsvp *rsvp.RSVP) error {
	db.Rsvp = rsvp
	return nil
}

func (db *MockSupabaseDatabase) GetRSVP(rsvpId string) (*rsvp.RSVP, error) {
	if db.Rsvp == nil {
		return nil, errors.New("No RSVP found")
	}
	return db.Rsvp, nil
}


func (db *MockSupabaseDatabase) GetAllRSVPs() ([]rsvp.RSVP, error) {
	if db.Rsvp == nil {
		return nil, errors.New("No RSVPs found")
	}
	return nil, nil
}
