package utils

import (
	"fmt"
	"hotelapi/models"
	"strconv"
	"strings"
)

func RoomIdGenerator(allrooms []models.Rooms) string {
	highestNumber := 0
	for _, room := range allrooms {
		numberString := strings.TrimPrefix(room.ID, "RM")
		currentNumber, err := strconv.Atoi(numberString)
		if err == nil && currentNumber > highestNumber {
			highestNumber = currentNumber
		}
	}
	return fmt.Sprintf("RM%03d", highestNumber+1)
}

func GuestIDGenerator(allguests []models.Guest) string {
	highestNumber := 0
	for _, guest := range allguests {
		numberString := strings.TrimPrefix(guest.ID, "GST")
		currentNumber, err := strconv.Atoi(numberString)
		if err == nil && currentNumber > highestNumber {
			highestNumber = currentNumber
		}
	}
	return fmt.Sprintf("GST%03d", highestNumber+1)
}

func BookingIDGenerator(allbookings []models.Booking) string {
	highestNumber := 0
	for _, booking := range allbookings {
		numberString := strings.TrimPrefix(booking.ID, "BK")
		currentNumber, err := strconv.Atoi(numberString)
		if err == nil && currentNumber > highestNumber {
			highestNumber = currentNumber
		}
	}
	return fmt.Sprintf("BK%03d", highestNumber+1)
}
