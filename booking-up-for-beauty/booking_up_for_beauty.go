package booking

import (
	"fmt"
	"time"
)

const dateHourLayout = "1/02/2006 15:04:05"
const layoutMediumDateAndTime = "January 2, 2006 15:04:05"
const layoutDayOfWeekLongDate = "Monday, January 2, 2006 15:04:05"
const layoutForDescription = "1/2/2006 15:04:05"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) (time.Time, error) {
	return time.Parse(dateHourLayout, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	givenDate, _ := time.Parse(layoutMediumDateAndTime, date)
	return time.Now().After(givenDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
//Thursday, May 13, 2010 20:32:00
func IsAfternoonAppointment(date string) bool {
	givenDate, _ := time.Parse(layoutDayOfWeekLongDate, date)
	hour := givenDate.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	givenDate, _ := time.Parse(layoutForDescription, date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", givenDate.Weekday(),
		givenDate.Month(), givenDate.Day(), givenDate.Year(), givenDate.Hour(), givenDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(
		time.Now().Year(),
		time.September,
		15,
		0,
		0,
		0,
		0,
		time.UTC,
	)
}
