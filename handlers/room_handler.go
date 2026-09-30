package handlers

import (
	"encoding/json"
	"hotelapi/models"
	"hotelapi/services"
	"hotelapi/storage"
	"hotelapi/utils"
	"net/http"
	"strconv"
)

var AllRooms []models.Rooms

func AddRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var newroom models.Rooms
		err := json.NewDecoder(r.Body).Decode(&newroom)
		if err != nil {
			http.Error(w, "Error getting request from user", http.StatusBadRequest)
			return
		}
		newroom.ID = utils.RoomIdGenerator(AllRooms)
		if services.DuplicationCheckOfRooms(AllRooms, newroom.RoomNumber) {
			http.Error(w, "This Room already Exists", http.StatusBadRequest)
			return
		}
		AllRooms = append(AllRooms, newroom)
		err = storage.SaveRooms(AllRooms)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Room Added Successfully!")

	}
}
func GetAllRooms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(AllRooms)
		if err != nil {
			http.Error(w, "Failed to get all rooms' data", http.StatusInternalServerError)
			return
		}
	}

}
func GetRoomByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		roomIdToFind := r.URL.Query().Get("id")
		if roomIdToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		foundRoom, found := services.GetRoomByID(AllRooms, roomIdToFind)
		if !found {
			http.Error(w, "Room Not Found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(foundRoom)

	}
}
func UpdateRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		roomIdToFind := r.URL.Query().Get("id")
		if roomIdToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		roomNumberToFind := r.URL.Query().Get("roomNumber")
		if roomNumberToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		roomNumber, err := strconv.Atoi(roomNumberToFind)
		if err != nil {
			http.Error(w, "Error accepting input", http.StatusBadRequest)
			return
		}
		_, found := services.SearchRoom(AllRooms, roomNumber, roomIdToFind)
		if !found {
			http.Error(w, "Room Not Found", http.StatusNotFound)
			return
		}
		var updateRoom models.Rooms
		err = json.NewDecoder(r.Body).Decode(&updateRoom)
		if err != nil {
			http.Error(w, "Error getting input from user!", http.StatusBadRequest)
			return
		}
		newupdatedRoom, status := services.UpdateRoom(roomNumber, AllRooms, updateRoom)
		if !status {
			http.Error(w, "Room Not Updated", http.StatusInternalServerError)
			return
		}
		AllRooms = newupdatedRoom
		//The reason for this is that updatedRoom is returning a full updated slice and not just a single room that has been updated
		err = storage.SaveRooms(AllRooms)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Room Updated Successfully!")
	}

}
func DeleteRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		roomIdToFind := r.URL.Query().Get("id")
		if roomIdToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		roomNumberToFind := r.URL.Query().Get("roomNumber")
		if roomNumberToFind == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		roomNumber, err := strconv.Atoi(roomNumberToFind)
		if err != nil {
			http.Error(w, "Error accepting input", http.StatusBadRequest)
			return
		}
		_, found := services.SearchRoom(AllRooms, roomNumber, roomIdToFind)
		if !found {
			http.Error(w, "Room Not Found", http.StatusNotFound)
			return
		}
		UpdatedRooms, status := services.DeleteRoom(AllRooms, roomNumber)
		if !status {
			http.Error(w, "Room Not Deleted", http.StatusInternalServerError)
			return
		}

		AllRooms = UpdatedRooms
		//The reason or this is that updatedRooms is returning a full updated slice minus the deleted room and not just a single room that has been updated
		err = storage.SaveRooms(AllRooms)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Room Deleted Successfully!")

	}
}
func GetAvailableRooms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		actualavailablerooms, status := services.GetAvailableRoom(AllRooms)
		if !status {
			http.Error(w, "No Available Room Found!", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(actualavailablerooms)

	}
}
func GetRoomByType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		roomTypeToGet := r.URL.Query().Get("type")
		if roomTypeToGet == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		roombytypefound, status := services.GetRoomByType(AllRooms, roomTypeToGet)
		if !status {
			http.Error(w, "No Room Found!", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(roombytypefound)

	}
}
func GetRoomByStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		roomStatusToGet := r.URL.Query().Get("status")
		if roomStatusToGet == "" {
			http.Error(w, "Missing required parameter", http.StatusBadRequest)
			return
		}
		roombystatusfound, status := services.GetRoomByStatus(AllRooms, roomStatusToGet)
		if !status {
			http.Error(w, "No Room Found!", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(roombystatusfound)

	}
}
