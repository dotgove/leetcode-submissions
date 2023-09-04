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

	// Consolidate setting zeros across rows and columns
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if idxRows[i] == 1 || idxCols[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
