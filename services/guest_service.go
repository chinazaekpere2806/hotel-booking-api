package services

import (
	//"fmt"
	"hotelapi/models"
	//"hotelapi/storage"
)

func DuplicationCheckOfGuests(allguests []models.Guest, guestphonenumber string) bool {
	for _, checks := range allguests {
		if checks.PhoneNumber == guestphonenumber {
			return true
		}
	}
	return false
}
func GetGuestByID(allguests []models.Guest, guestId string) (models.Guest, bool) {
	for _, findroom := range allguests {
		if findroom.ID == guestId {
			return findroom, true
		}
	}
	return models.Guest{}, false
}
func SearchGuest(allguests []models.Guest, guestId string) (models.Guest, bool) {
	for _, guests := range allguests {
		if guests.ID == guestId {
			return guests, true
		}
	}
	return models.Guest{}, false
}
func UpdateGuest(allguests []models.Guest, updatedguests models.Guest, guestId string) ([]models.Guest, bool) {
	found := false
	for i, currentguest := range allguests {
		if currentguest.ID == guestId {
			allguests[i].Firstname = updatedguests.Firstname
			allguests[i].Lastname = updatedguests.Lastname
			allguests[i].PhoneNumber = updatedguests.PhoneNumber
			allguests[i].Email = updatedguests.Email
			found = true
			break
		}
	}
	return allguests, found
}
func DeleteGuests(guests []models.Guest, guestId string) ([]models.Guest, bool) {
	foundIndex := -1
	for i, guest := range guests {
		if guestId == guest.ID {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return guests, false
	}
	updatedGuests := append(guests[:foundIndex], guests[foundIndex+1:]...)
	return updatedGuests, true

}
