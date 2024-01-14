package rsvp

type RSVP struct {
	RsvpId string `json:"rsvpId" validate:"required"`
	RsvpName string `json:"rsvpName"`
	NumGuests int `json:"numGuests" validate:"required"`
	PlusOneName string `json:"plusOneName"`
	MealSelection1 string `json:"mealSelection1" validate:"required"`
	MealSelection2 string `json:"mealSelection2"`
	ExtraText string `json:"extraText"`
	Attending bool `json:"attending" validate:"required"`
	Added string `json:"added"`
}