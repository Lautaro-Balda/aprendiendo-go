package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	time, _ := time.Parse("1/02/2006 15:04:05", date)
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	ahora := time.Now()
	fecha, _ := time.Parse("January 2, 2006 15:04:05", date)
	return fecha.Before(ahora)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	fecha, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if fecha.Hour() >= 12 && fecha.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	fecha, _ := time.Parse("1/2/2006 15:04:05", date)
	return fecha.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	openDay := "2012-09-15"
	currentYear := fmt.Sprintf("%d", time.Now().Year())
	formattedTime := currentYear + openDay[4:]
	parsedDate, _ := time.Parse("2006-01-02", formattedTime)
	return parsedDate

}
