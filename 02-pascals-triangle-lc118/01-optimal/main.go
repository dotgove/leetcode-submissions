package main

/*
Stats
	Runtime: 1 ms, faster than 90.55% of Go online submissions for Pascal's Triangle.
	Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Pascal's Triangle.
*/

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}
