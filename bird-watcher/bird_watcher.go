package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, value := range birdsPerDay {
		count += value
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	index := (week * 7) - 7
	count := 0
	for daysCount := 0; daysCount < 7; daysCount++ {
		count += birdsPerDay[index]
		index++
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index := 0; index < len(birdsPerDay); index += 2 {
		birdsPerDay[index] += 1
	}
	return birdsPerDay
}
