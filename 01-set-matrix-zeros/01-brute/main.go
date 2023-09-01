package main

func setZeroes(matrix [][]int) {
	lCol := len(matrix)
	lRow := len(matrix[0])

	setZerosInRow := func(rowIdx int) {
		for j := 0; j < lRow; j++ {
			if matrix[rowIdx][j] > 0 {
				matrix[rowIdx][j] = -1
			}
		}
	}

	setZerosInColumn := func(colIdx int) {
		for i := 0; i < lCol; i++ {
			if matrix[i][colIdx] > 0 {
				matrix[i][colIdx] = -1
			}
		}
	}

	// Convert all non zero values to -1
	for i := 0; i < lCol; i++ {
		for j := 0; j < lRow; j++ {
			if matrix[i][j] == 0 {
				setZerosInRow(i)
				setZerosInColumn(j)
			}
		}
	}

	// Convert all -1 values to 0
	for i := 0; i < lCol; i++ {
		for j := 0; j < lRow; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}
