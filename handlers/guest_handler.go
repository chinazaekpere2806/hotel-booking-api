package handlers

import (
	"encoding/json"
	"hotelapi/models"
	"hotelapi/services"
	"hotelapi/storage"
	"hotelapi/utils"
	"net/http"
)

var Allguests []models.Guest

func CreateGuest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var newguest models.Guest
		err := json.NewDecoder(r.Body).Decode(&newguest)
		if err != nil {
			http.Error(w, "Error accepting input!", http.StatusBadRequest)
			return
		}
		newguest.ID = utils.GuestIDGenerator(Allguests)
		if services.DuplicationCheckOfGuests(Allguests, newguest.PhoneNumber) {
			http.Error(w, "A guest with this phone number already exists", http.StatusConflict)
			return
		}
		if !utils.IsValidSyntax(newguest.Email) {
			http.Error(w, "Email Inputted Is Not Valid!", http.StatusBadRequest)
			return
		}
		if !utils.HasMailServer(newguest.Email) {
			http.Error(w, "Email Doesn't Have A Server!", http.StatusBadRequest)
			return
		}
		Allguests = append(Allguests, newguest)
		err = storage.SaveGuests(Allguests)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Guest Data Created Successfully!")

	}
}
func GetAllGuests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(Allguests)
		if err != nil {
			http.Error(w, "Failed to get all guests' data", http.StatusInternalServerError)
			return
		}
	}
}
func GetGuestById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		GuestIdToFind := r.URL.Query().Get("id")
		if GuestIdToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		GuestIdFound, status := services.GetGuestByID(Allguests, GuestIdToFind)
		if !status {
			http.Error(w, "Guest Not Found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(GuestIdFound)
	}
}
func UpdateGuest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		guestIdToFind := r.URL.Query().Get("id")
		if guestIdToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		_, status := services.SearchGuest(Allguests, guestIdToFind)
		if !status {
			http.Error(w, "Guest Data isn't found!", http.StatusNotFound)
			return
		}
		var updatedguestdetails models.Guest
		err := json.NewDecoder(r.Body).Decode(&updatedguestdetails)
		if err != nil {
			http.Error(w, "Error getting updated guest details!", http.StatusBadRequest)
			return
		}
		updatedguestslice, statusofupdate := services.UpdateGuest(Allguests, updatedguestdetails, guestIdToFind)
		if !statusofupdate {
			http.Error(w, "Error updating guest details!", http.StatusInternalServerError)
			return
		}
		Allguests = updatedguestslice
		err = storage.SaveGuests(Allguests)
		if err != nil {
			http.Error(w, "Error saving updated guests list!", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Guest data has been updated!")

	}
}
func DeleteGuest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		guestIdToFind := r.URL.Query().Get("id")
		if guestIdToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		_, status := services.SearchGuest(Allguests, guestIdToFind)
		if !status {
			http.Error(w, "Guest Data isn't found!", http.StatusNotFound)
			return
		}
		UpdatedGuestslice, statusofdelete := services.DeleteGuests(Allguests, guestIdToFind)
		if !statusofdelete {
			http.Error(w, "Guest data couldn't be deleted!", http.StatusNotFound)
			return
		}
		Allguests = UpdatedGuestslice
		err := storage.SaveGuests(Allguests)
		if err != nil {
			http.Error(w, "Error saving updated guest list", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Guest data deleted sucessfully!")
	}
}
