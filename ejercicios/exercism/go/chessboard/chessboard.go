package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	files, existe := cb[file]
	if !existe {
		return 0
	}
	var cuadradosOcupados int

	for _, ocupado := range files {
		if ocupado {
			cuadradosOcupados++
		}
	}

	return cuadradosOcupados

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	var cuadradosOcupados int
	for _, columna := range cb {
		if columna[rank-1] {
			cuadradosOcupados++
		}
	}

	return cuadradosOcupados
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var cuadrados int
	for _, value := range cb {
		for i := 0; i < len(value); i++ {
			cuadrados++
		}
	}
	return cuadrados
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var cuadradosOcupados int
	for _, value := range cb {
		for _, value := range value {
			if value {
				cuadradosOcupados++
			}
		}
	}
	return cuadradosOcupados
}
