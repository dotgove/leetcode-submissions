package main

func setZeroes(matrix [][]int) {
	lCol := len(matrix)
	lRow := len(matrix[0])

	// Sets non zero (>0) values in a given row to -1
	markRowValues := func(rowIdx int) {
		for j := 0; j < lRow; j++ {
			if matrix[rowIdx][j] > 0 {
				matrix[rowIdx][j] = -1
			}
		}
	}

	// Sets non zero (>0) values in a given column to -1
	markColumnValues := func(colIdx int) {
		for i := 0; i < lCol; i++ {
			if matrix[i][colIdx] > 0 {
				matrix[i][colIdx] = -1
			}
		}
	}

	// Mark all non zero values to -1
	for i := 0; i < lCol; i++ {
		for j := 0; j < lRow; j++ {
			if matrix[i][j] == 0 {
				markRowValues(i)
				markColumnValues(j)
			}
		}
	}

	// Mark all -1 values to 0
	for i := 0; i < lCol; i++ {
		for j := 0; j < lRow; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}
