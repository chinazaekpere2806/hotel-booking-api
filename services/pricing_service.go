package services

import "hotelapi/models"

func CalculateTotalCost(room models.Rooms, duration int, numberOfGuests int) (float64, bool) {
	if duration <= 0 || numberOfGuests <= 0 || numberOfGuests > room.Capacity || room.PricePerNight < 0 {
		return 0, false
	}
	return room.PricePerNight * float64(duration), true
}
