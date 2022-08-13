package logs

import "unicode/utf8"

var applications = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		name, exists := applications[char]
		if exists {
			return name
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for index, char := range log {
		if char == oldRune {
			runes[index] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
