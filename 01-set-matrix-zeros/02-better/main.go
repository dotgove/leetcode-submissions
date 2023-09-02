package main

func setZeroes(matrix [][]int) {
	nRows := len(matrix)
	nCols := len(matrix[0])

	// Storing marked indexes
	idxRows := make([]int, nRows)
	idxCols := make([]int, nCols)

	// Identify and mark indexes of cols and rows to be set to 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if matrix[i][j] == 0 {
				idxCols[j] = 1
				idxRows[i] = 1
			}
		}
	}

	// Set the respective rows in matrix to 0
	for i := 0; i < len(idxRows); i++ {
		if idxRows[i] == 0 {
			continue
		}
		for j := 0; j < nCols; j++ {
			matrix[i][j] = 0
		}
	}

	// Set the respective cols in matrix to 0
	for j := 0; j < nCols; j++ {
		if idxCols[j] == 0 {
			continue
		}
		for i := 0; i < nRows; i++ {
			matrix[i][j] = 0
		}
	}
}

// func main() {
// 	input := [][]int{
// 		{0, 1, 2, 0},
// 		{3, 4, 5, 2},
// 		{1, 3, 1, 5},
// 	}

// 	setZeroes(input)
// 	for _, row := range input {
// 		fmt.Println(row)
// 	}
// }
