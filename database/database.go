package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mittonface/wedding-backend/rsvp"
	"github.com/supabase-community/supabase-go"
	"github.com/supabase/postgrest-go"
)

type SupabaseDatabase struct {
	Client *supabase.Client
 }

type RsvpDB interface {
	Initialize() error
	InsertRSVP(rsvp *rsvp.RSVP) error
	GetRSVP(rsvpId string) (*rsvp.RSVP, error)
	GetAllRSVPs() ([]rsvp.RSVP, error) // New method
	DeleteRSVPs(rsvpId string) error

}

func (db *SupabaseDatabase) Initialize() error {
	err := godotenv.Load(".env")
	SUPABASE_URL := os.Getenv("SUPABASE_URL")
	SUPABASE_SECRET := os.Getenv("SUPABASE_SECRET")
	client, err := supabase.NewClient(SUPABASE_URL, SUPABASE_SECRET, nil)
	if err != nil {
		return err
	}
	db.Client = client
	return nil
}

func (db *SupabaseDatabase) InsertRSVP(rsvp *rsvp.RSVP) error {
	_, _, err := db.Client.From("rsvps").Insert(rsvp, false, "", "", "").Execute()
	if err != nil {
		return err
	}
	return nil
}

func (db *SupabaseDatabase) GetRSVP(rsvpId string) (*rsvp.RSVP, error) {

	var rsvpResults []rsvp.RSVP
	_, err := db.Client.From("rsvps").Select("*", "exact", false).Eq("rsvpId", rsvpId).Order("added", &postgrest.OrderOpts{Ascending: false, NullsFirst: false, ForeignTable: ""}).Limit(1, "").ExecuteTo(&rsvpResults)
	// Unmarshal the bytes into the rsvp struct
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling result into RSVP struct: %w", err)
	}

	return &rsvpResults[0], nil
 }

 func (db *SupabaseDatabase) GetAllRSVPs() ([]rsvp.RSVP, error) {
	var rsvpResults []rsvp.RSVP
	_, err := db.Client.From("rsvps").Select("*", "exact", false).Neq("rsvpId", "dummy").
	ExecuteTo(&rsvpResults)
	if err != nil {
		return nil, fmt.Errorf("error retrieving all RSVPs: %w", err)
	}
	return rsvpResults, nil
}

func (db *SupabaseDatabase) DeleteRSVPs(rsvpId string) error {
	_, _, err := db.Client.From("rsvps").Delete("", "exact").Eq("rsvpId", rsvpId).Execute()
	if err != nil {
		return fmt.Errorf("error retrieving all RSVPs: %w", err)
	}
	return nil
}