package chessboard

// File Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Chessboard Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	row, exists := cb[file]
	count := 0
	if exists {
		for _, filled := range row {
			if filled {
				count++
			}
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank >= 1 && rank <= 8 {
		for _, file := range cb {
			if file[rank-1] {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for range file {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, present := range file {
			if present {
				count++
			}
		}
	}
	return count
}
