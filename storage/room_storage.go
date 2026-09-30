package storage

import (
	"encoding/json"
	"hotelapi/models"
	"os"
)

func LoadRooms() ([]models.Rooms, error) {
	jsondata, err := os.ReadFile("rooms.json")
	if err != nil {
		return nil, err
	}
	var r []models.Rooms
	err = json.Unmarshal(jsondata, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func SaveRooms(r []models.Rooms) error {
	savedrooms, err := json.MarshalIndent(r, " ", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("rooms.json", savedrooms, 0664)
	if err != nil {
		return err
	}
	return nil
}
