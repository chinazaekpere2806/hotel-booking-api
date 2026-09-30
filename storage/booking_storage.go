package storage

import (
	"encoding/json"
	"hotelapi/models"
	"os"
)

func LoadBookings() ([]models.Booking, error) {
	jsondata, err := os.ReadFile("bookings.json")
	if err != nil {
		return nil, err
	}
	var b []models.Booking
	err = json.Unmarshal(jsondata, &b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func SaveBookings(b []models.Booking) error {
	savedbookings, err := json.MarshalIndent(b, " ", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("bookings.json", savedbookings, 0664)
	if err != nil {
		return err
	}
	return nil
}
