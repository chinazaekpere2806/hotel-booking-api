package handlers

import (
	"encoding/json"
	"hotelapi/models"
	"hotelapi/services"
	"hotelapi/storage"
	"hotelapi/utils"
	"net/http"
)

var Allbookings []models.Booking

func CreateBooking() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var request models.Booking
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if request.Guestid.ID == "" || request.Roomid.ID == "" {
			http.Error(w, "Guest ID and Room ID are required", http.StatusBadRequest)
			return
		}
		if request.Duration <= 0 || request.NumberOfGuests <= 0 {
			http.Error(w, "Duration and number of guests must be greater than zero", http.StatusBadRequest)
			return
		}

		guest, found := services.SearchGuest(Allguests, request.Guestid.ID)
		if !found {
			http.Error(w, "Guest Data Doesn't Exist, So Booking Can't Be Created!", http.StatusNotFound)
			return
		}

		room, found := services.SearchRoomforBooking(AllRooms, request.Roomid.ID)
		if !found {
			http.Error(w, "Room Data couldn't be found!", http.StatusNotFound)
			return
		}
		if room.Status != "Available" {
			http.Error(w, "The Room is not available for booking!", http.StatusConflict)
			return
		}
		if request.NumberOfGuests > room.Capacity {
			http.Error(w, "Number of guests exceeds room capacity", http.StatusBadRequest)
			return
		}

		totalCost, valid := services.CalculateTotalCost(room, request.Duration, request.NumberOfGuests)
		if !valid {
			http.Error(w, "Invalid booking details", http.StatusBadRequest)
			return
		}

		request.ID = utils.BookingIDGenerator(Allbookings)
		request.Guestid = guest
		request.Roomid = room
		request.RoomType = room
		request.Checkout = utils.CheckoutDate(request.Checkin, request.Duration)
		request.TotalCost = totalCost
		request.Status = "Confirmed"

		if !services.MarkRoomAsBooked(AllRooms, room.ID) {
			http.Error(w, "The Room could not be booked", http.StatusConflict)
			return
		}
		if err := storage.SaveRooms(AllRooms); err != nil {
			// Roll back the in-memory room state if persistence fails.
			services.MarkRoomAsAvailable(AllRooms, room.ID)
			http.Error(w, "Error saving room data", http.StatusInternalServerError)
			return
		}

		Allbookings = append(Allbookings, request)
		if err := storage.SaveBookings(Allbookings); err != nil {
			// Roll back both in-memory state and the room file.
			Allbookings = Allbookings[:len(Allbookings)-1]
			services.MarkRoomAsAvailable(AllRooms, room.ID)
			_ = storage.SaveRooms(AllRooms)
			http.Error(w, "Error saving booking data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(request)
	}
}

func GetBookingById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		bookingID := r.URL.Query().Get("id")
		if bookingID == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		booking, found := services.GetBookingById(Allbookings, bookingID)
		if !found {
			http.Error(w, "The Requested Booking Couldn't Be Found!", http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(w).Encode(booking)
	}
}

func GetAllBookings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		_ = json.NewEncoder(w).Encode(Allbookings)
	}
}

func GetBookingByGuest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		guestID := r.URL.Query().Get("guestId")
		if guestID == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		booking, found := services.GetBookingByGuest(Allbookings, guestID)
		if !found {
			http.Error(w, "The booking couldn't be found!", http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(w).Encode(booking)
	}
}

func CancelBooking() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodDelete && r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		bookingID := r.URL.Query().Get("id")
		guestID := r.URL.Query().Get("guestId")
		roomID := r.URL.Query().Get("roomId")
		if bookingID == "" || guestID == "" || roomID == "" {
			http.Error(w, "id, guestId and roomId are required", http.StatusBadRequest)
			return
		}

		booking, found := services.GetBookingById(Allbookings, bookingID)
		if !found || booking.Guestid.ID != guestID || booking.Roomid.ID != roomID {
			http.Error(w, "The Requested Booking Couldn't Be Found!", http.StatusNotFound)
			return
		}
		if !services.CancelBooking(Allbookings, bookingID, guestID) {
			http.Error(w, "Only confirmed bookings can be cancelled", http.StatusConflict)
			return
		}
		if !services.ReleaseBookedRoom(AllRooms, roomID) {
			for i := range Allbookings {
				if Allbookings[i].ID == bookingID {
					Allbookings[i].Status = "Confirmed"
				}
			}
			http.Error(w, "Room is not in Booked state", http.StatusConflict)
			return
		}
		if err := storage.SaveBookings(Allbookings); err != nil {
			services.MarkRoomAsBooked(AllRooms, roomID)
			_ = storage.SaveRooms(AllRooms)
			http.Error(w, "Error saving booking data", http.StatusInternalServerError)
			return
		}
		if err := storage.SaveRooms(AllRooms); err != nil {
			http.Error(w, "Error saving room data", http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode("Booking Has Been Cancelled!")
	}
}
func CheckIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		bookingID, guestID, roomID := r.URL.Query().Get("id"), r.URL.Query().Get("guestId"), r.URL.Query().Get("roomId")
		if bookingID == "" || guestID == "" || roomID == "" {
			http.Error(w, "id, guestId and roomId are required", http.StatusBadRequest)
			return
		}
		booking, found := services.GetBookingById(Allbookings, bookingID)
		if !found || booking.Guestid.ID != guestID || booking.Roomid.ID != roomID {
			http.Error(w, "The Requested Booking Couldn't Be Found!", http.StatusNotFound)
			return
		}
		if !services.CheckInForGuest(Allbookings, bookingID, guestID) {
			http.Error(w, "Only confirmed bookings can be checked in", http.StatusConflict)
			return
		}
		if !services.MarkRoomAsOccupied(AllRooms, roomID) {
			for i := range Allbookings {
				if Allbookings[i].ID == bookingID {
					Allbookings[i].Status = "Confirmed"
				}
			}
			http.Error(w, "Room is not in Booked state", http.StatusConflict)
			return
		}
		if err := storage.SaveBookings(Allbookings); err != nil {
			services.MarkRoomAsBooked(AllRooms, roomID)
			_ = storage.SaveRooms(AllRooms)
			http.Error(w, "Error saving booking data", http.StatusInternalServerError)
			return
		}
		if err := storage.SaveRooms(AllRooms); err != nil {
			http.Error(w, "Error saving room data", http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode("Check-In Successful!")
	}
}

func CheckOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		bookingID, guestID, roomID := r.URL.Query().Get("id"), r.URL.Query().Get("guestId"), r.URL.Query().Get("roomId")
		if bookingID == "" || guestID == "" || roomID == "" {
			http.Error(w, "id, guestId and roomId are required", http.StatusBadRequest)
			return
		}
		booking, found := services.GetBookingById(Allbookings, bookingID)
		if !found || booking.Guestid.ID != guestID || booking.Roomid.ID != roomID {
			http.Error(w, "The Requested Booking Couldn't Be Found!", http.StatusNotFound)
			return
		}
		if !services.CheckOutForGuest(Allbookings, bookingID, guestID) {
			http.Error(w, "Only checked-in bookings can be checked out", http.StatusConflict)
			return
		}
		if !services.MarkRoomAsCleaning(AllRooms, roomID) {
			for i := range Allbookings {
				if Allbookings[i].ID == bookingID {
					Allbookings[i].Status = "Checked-In"
				}
			}
			http.Error(w, "Room is not in Occupied state", http.StatusConflict)
			return
		}
		if !services.MarkRoomAsAvailable(AllRooms, roomID) {
			services.MarkRoomAsOccupied(AllRooms, roomID)
			for i := range Allbookings {
				if Allbookings[i].ID == bookingID {
					Allbookings[i].Status = "Checked-In"
				}
			}
			http.Error(w, "Room could not be returned to Available state", http.StatusConflict)
			return
		}
		if err := storage.SaveBookings(Allbookings); err != nil {
			services.MarkRoomAsOccupied(AllRooms, roomID)
			_ = storage.SaveRooms(AllRooms)
			http.Error(w, "Error saving booking data", http.StatusInternalServerError)
			return
		}
		if err := storage.SaveRooms(AllRooms); err != nil {
			http.Error(w, "Error saving room data", http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode("Check-Out Successful!")
	}
}
