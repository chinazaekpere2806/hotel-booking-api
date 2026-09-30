package services

import "hotelapi/models"

func SearchRoomforBooking(rooms []models.Rooms, roomID string) (models.Rooms, bool) {
	for _, room := range rooms {
		if room.ID == roomID {
			return room, true
		}
	}
	return models.Rooms{}, false
}

func CheckRoomAvailabilty(rooms []models.Rooms, roomID string) bool {
	room, found := SearchRoomforBooking(rooms, roomID)
	return found && room.Status == "Available"
}

func MarkRoomAsBooked(rooms []models.Rooms, roomID string) bool {
	for i := range rooms {
		if rooms[i].ID == roomID && rooms[i].Status == "Available" {
			rooms[i].Status = "Booked"
			return true
		}
	}
	return false
}

func GetBookingById(bookings []models.Booking, bookingID string) (models.Booking, bool) {
	for _, booking := range bookings {
		if booking.ID == bookingID {
			return booking, true
		}
	}
	return models.Booking{}, false
}

func GetBookingByGuest(bookings []models.Booking, guestID string) (models.Booking, bool) {
	for _, booking := range bookings {
		if booking.Guestid.ID == guestID {
			return booking, true
		}
	}
	return models.Booking{}, false
}

func CancelBooking(bookings []models.Booking, bookingID, guestID string) bool {
	for i := range bookings {
		if bookings[i].ID == bookingID && bookings[i].Guestid.ID == guestID && bookings[i].Status == "Confirmed" {
			bookings[i].Status = "Cancelled"
			return true
		}
	}
	return false
}

func CheckBookingTrue(bookings []models.Booking, bookingID, guestID string) bool {
	for _, booking := range bookings {
		if booking.ID == bookingID && booking.Guestid.ID == guestID {
			return true
		}
	}
	return false
}

func CheckInForGuest(bookings []models.Booking, bookingID, guestID string) bool {
	for i := range bookings {
		if bookings[i].ID == bookingID && bookings[i].Guestid.ID == guestID && bookings[i].Status == "Confirmed" {
			bookings[i].Status = "Checked-In"
			return true
		}
	}
	return false
}

func MarkRoomAsOccupied(rooms []models.Rooms, roomID string) bool {
	for i := range rooms {
		if rooms[i].ID == roomID && rooms[i].Status == "Booked" {
			rooms[i].Status = "Occupied"
			return true
		}
	}
	return false
}

func CheckOutForGuest(bookings []models.Booking, bookingID, guestID string) bool {
	for i := range bookings {
		if bookings[i].ID == bookingID && bookings[i].Guestid.ID == guestID && bookings[i].Status == "Checked-In" {
			bookings[i].Status = "Checked-Out"
			return true
		}
	}
	return false
}

func MarkRoomAsCleaning(rooms []models.Rooms, roomID string) bool {
	for i := range rooms {
		if rooms[i].ID == roomID && rooms[i].Status == "Occupied" {
			rooms[i].Status = "Cleaning"
			return true
		}
	}
	return false
}

func ReleaseBookedRoom(rooms []models.Rooms, roomID string) bool {
	for i := range rooms {
		if rooms[i].ID == roomID && rooms[i].Status == "Booked" {
			rooms[i].Status = "Available"
			return true
		}
	}
	return false
}

func MarkRoomAsAvailable(rooms []models.Rooms, roomID string) bool {
	for i := range rooms {
		if rooms[i].ID == roomID && rooms[i].Status == "Cleaning" {
			rooms[i].Status = "Available"
			return true
		}
	}
	return false
}
