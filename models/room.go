package models

type Rooms struct {
	ID            string  `json:"id"`
	RoomNumber    int     `json:"roomNumber"`
	Type          string  `json:"type"`
	Capacity      int     `json:"capacity"`
	PricePerNight float64 `json:"pricePerNight"`
	Status        string  `json:"status"`
	Description   string  `json:"description"`
}
