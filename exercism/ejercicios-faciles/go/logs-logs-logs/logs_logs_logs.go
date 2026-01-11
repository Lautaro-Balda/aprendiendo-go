package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, value := range log {
		runaConvertida := string(value)
		switch runaConvertida {
		case "❗":
			return "recommendation"
		case "🔍":
			return "search"
		case "☀":
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for i, rune := range runes {
		if rune == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	tamañoEnRunas := utf8.RuneCountInString(log)
	return tamañoEnRunas <= limit
}
