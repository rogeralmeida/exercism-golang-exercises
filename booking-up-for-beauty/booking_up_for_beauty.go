package booking

import (
	"fmt"
	"time"
)

const dateHourLayout = "1/2/2006 15:04:05"
const layoutMediumDateAndTime = "January 2, 2006 15:04:05"
const layoutDayOfWeekLongDate = "Monday, January 2, 2006 15:04:05"
const layoutForDescription = "1/2/2006 15:04:05"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	givenDate, _ := time.Parse(dateHourLayout, date)
	return givenDate
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
	givenDate := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s.", givenDate.Weekday(),
		givenDate.Format("January 2, 2006, at 15:04"))
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
