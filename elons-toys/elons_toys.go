package elon

import "fmt"

func (car *Car) Drive() {
	if car.CanFinish(car.speed) {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%s", car.battery, "%")
}

func (car Car) CanFinish(trackDistance int) bool {
	laps := trackDistance / car.speed
	return laps*car.batteryDrain <= car.battery
}
