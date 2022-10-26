package utils

import "encoding/json"

/*This file have the matrix sequence logic implemented*/

// Horizontal: Horizontal function is to read the matrix in the horizontal direction and see the sequences inside the matrix
func Horizontal(letters []string, length int64) (total int64) {
	for _, letter := range letters {
		var count int64

		for j := 0; j < len(letter)-1; j++ {
			if letter[j] == letter[j+1] {
				count++
			} else {
				count = 0
			}

			if count == length-1 {
				total++
			}
		}
	}

	return
}

// Vertical: Vertical function is to read the matrix in the vertical direction and see the sequences inside the matrix
func Vertical(letters []string, length int64) (total int64) {
	for i := 0; i < len(letters[0]); i++ {
		var count int64

		for j := 0; j < len(letters)-1; j++ {
			if letters[j][i] == letters[j+1][i] {
				count++
			} else {
				count = 0
			}

			if count == length-1 {
				total++
			}
		}
	}

	return
}

// BottomUpOblique: This function it's to read the matrix to the bottom for up in diagonal in the L > R direction
func BottomUpOblique(letters []string, length int64) int64 {
	var min func(int, int) int
	var max func(int, int) int

	min = func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	max = func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	rows := len(letters)
	cols := len(letters[0])

	var aux []string

	for i := 1; i <= rows+cols-1; i++ {
		start := max(0, i-rows)
		count := min(min(i, cols-start), rows)

		for j := 0; j < count; j++ {
			aux = append(aux, string(letters[min(rows, i)-j-1][start+j]))
		}
	}

	return Count(aux, length)
}

// UpBottomOblique: I used recursive logic for read the matrix to the up for bottom in diagonal in the R > L direction
func UpBottomOblique(letters []string, length int64) int64 {
	var w, k int
	var flag bool = true
	var aux []string
	var readUpBottom func([]string, int, int, int, int) bool

	readUpBottom = func(letters []string, i, j, rows, cols int) bool {
		if i >= rows || j >= cols {
			a := w
			w = k
			k = a

			if flag {
				flag = !flag
				w++
			} else {
				flag = !flag
			}

			return false
		}

		aux = append(aux, string(letters[i][j]))

		if readUpBottom(letters, i+1, j+1, rows, cols) {
			return true
		}

		if readUpBottom(letters, w, k, rows, cols) {
			return true
		}

		return true
	}

	readUpBottom(letters, 0, 0, len(letters), len(letters[0]))

	return Count(aux, length)
}

// Count: For the Oblique I used this function to count the total of valids sequences inside the array
func Count(array []string, length int64) (total int64) {
	var count int64

	for i, letter := range array {
		if i < len(array)-1 {
			if letter == array[i+1] {
				count++
			} else {
				count = 0
			}

			if count == length-1 {
				total++
			}
		}
	}

	return
}

// ValidateLetters: This function it's to check if inside the sequence has just the [B, D H, U] letters
func ValidateLetters(letters []string) bool {
	isValid := true

	for _, letter := range letters {
		for _, l := range letter {
			if l == 'B' || l == 'D' || l == 'H' || l == 'U' {
				isValid = true
			} else {
				isValid = false
				break
			}
		}

		if !isValid {
			break
		}
	}

	return isValid
}

// SequenceProcessValidate: This function it's to check if the sequence satisfy the conditions
func SequenceProcessValidate(letters []string, consecutive_length int64) (bool, string) {
	// Count the sequences in the matrix
	total := (Horizontal(letters, consecutive_length) +
		Vertical(letters, consecutive_length) +
		BottomUpOblique(letters, consecutive_length) +
		UpBottomOblique(letters, consecutive_length))

	// Consulting the letters inside the sequences
	isValid := ValidateLetters(letters) && total >= 2

	// Marshal the letters and stringfy
	lMarshal, _ := json.Marshal(letters)

	return isValid, string(lMarshal)
}
