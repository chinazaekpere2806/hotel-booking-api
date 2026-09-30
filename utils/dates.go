package utils

import (
	"time"
)

func CheckoutDate(checkinDate time.Time, stayduration int) time.Time {
	checkOutDate := checkinDate.AddDate(0, 0, stayduration)
	return checkOutDate
}
