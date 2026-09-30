package storage

import (
	"encoding/json"
	"hotelapi/models"
	"os"
)

func LoadGuests() ([]models.Guest, error) {
	jsondata, err := os.ReadFile("guests.json")
	if err != nil {
		return nil, err
	}
	var g []models.Guest
	err = json.Unmarshal(jsondata, &g)
	if err != nil {
		return nil, err
	}
	return g, nil
}
func SaveGuests(g []models.Guest) error {
	savedguests, err := json.MarshalIndent(g, " ", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("guests.json", savedguests, 0664)
	if err != nil {
		return err
	}
	return nil
}
