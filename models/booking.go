package models

import "time"

type Booking struct {
	ID             string    `json:"id"`
	Guestid        Guest     `json:"guestId"`
	Roomid         Rooms     `json:"roomId"`
	RoomType       Rooms     `json:"roomType"`
	Checkin        time.Time `json:"checkIn"`
	Duration       int       `json:"duration"`
	Checkout       time.Time `json:"checkOut"`
	NumberOfGuests int       `json:"numberOfGuests"`
	Status         string    `json:"status"`
	TotalCost      float64   `json:"totalCost"`
}
