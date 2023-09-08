package setmatrixzeroslc73

func setZeroes(matrix [][]int) {
	var hasFirstRowZero, hasFirstColZero bool

	// Identify zeros within the matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					hasFirstRowZero = true
				}
				if j == 0 {
					hasFirstColZero = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set respective columns and rows as zero
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// If the first column of the matrix has a zero, Set zeros across the first column
	if hasFirstColZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	// If the first row of the matrix has a zero, Set zeros across the first row
	if hasFirstRowZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
}
