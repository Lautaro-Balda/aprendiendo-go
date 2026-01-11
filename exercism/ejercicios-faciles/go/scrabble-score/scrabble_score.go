package scrabble

func Score(word string) int {
	var points int
	for _, value := range word {
		switch string(value) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T", "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			points += 1
		case "D", "G", "d", "g":
			points += 2
		case "B", "C", "M", "P", "b", "c", "m", "p":
			points += 3
		case "F", "H", "V", "W", "Y", "f", "h", "v", "w", "y":
			points += 4
		case "K", "k":
			points += 5
		case "J", "X", "j", "x":
			points += 8
		case "Q", "Z", "q", "z":
			points += 10
		}
	}
	return points
}
