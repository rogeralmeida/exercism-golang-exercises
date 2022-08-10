package purchase

import (
	"fmt"
)

const fifthPercentage = 0.5
const eightPercentage = 0.8
const seventhPercentage = 0.7

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	winnerOption := option2
	if option1 < option2 {
		winnerOption = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", winnerOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	percentage := fifthPercentage
	if age <= 3 {
		percentage = eightPercentage
	} else if age < 9 {
		percentage = seventhPercentage
	}
	return originalPrice * percentage
}
