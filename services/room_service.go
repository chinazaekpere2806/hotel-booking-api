package services

import "hotelapi/models"

func DuplicationCheckOfRooms(rooms []models.Rooms, roomNumber int) bool {
	for _, room := range rooms {
		if room.RoomNumber == roomNumber {
			return true
		}
	}
	return false
}

func SearchRoom(rooms []models.Rooms, roomNumber int, roomID string) (models.Rooms, bool) {
	for _, room := range rooms {
		if room.RoomNumber == roomNumber && room.ID == roomID {
			return room, true
		}
	}
	return models.Rooms{}, false
}

func UpdateRoom(roomNumber int, rooms []models.Rooms, updatedRoom models.Rooms) ([]models.Rooms, bool) {
	for i := range rooms {
		if rooms[i].RoomNumber == roomNumber {
			rooms[i].Capacity = updatedRoom.Capacity
			rooms[i].Description = updatedRoom.Description
			rooms[i].Status = updatedRoom.Status
			rooms[i].PricePerNight = updatedRoom.PricePerNight
			rooms[i].Type = updatedRoom.Type
			return rooms, true
		}
	}
	return rooms, false
}

func DeleteRoom(rooms []models.Rooms, roomNumber int) ([]models.Rooms, bool) {
	for i, room := range rooms {
		if room.RoomNumber == roomNumber {
			return append(rooms[:i], rooms[i+1:]...), true
		}
	}
	return rooms, false
}

func GetRoomByID(rooms []models.Rooms, roomID string) (models.Rooms, bool) {
	for _, room := range rooms {
		if room.ID == roomID {
			return room, true
		}
	}
	return models.Rooms{}, false
}

func GetAvailableRoom(rooms []models.Rooms) ([]models.Rooms, bool) {
	var availableRooms []models.Rooms
	for _, room := range rooms {
		if room.Status == "Available" {
			availableRooms = append(availableRooms, room)
		}
	}
	return availableRooms, len(availableRooms) > 0
}

func GetRoomByType(rooms []models.Rooms, roomType string) ([]models.Rooms, bool) {
	var matches []models.Rooms
	for _, room := range rooms {
		if room.Type == roomType {
			matches = append(matches, room)
		}
	}
	return matches, len(matches) > 0
}

func GetRoomByStatus(rooms []models.Rooms, roomStatus string) ([]models.Rooms, bool) {
	var matches []models.Rooms
	for _, room := range rooms {
		if room.Status == roomStatus {
			matches = append(matches, room)
		}
	}
	return matches, len(matches) > 0
}
