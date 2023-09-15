package maxsubarraylc53

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Example 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"Example 2", []int{1}, 1},
		{"Example 3", []int{5, 4, -1, 7, 8}, 23},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := maxSubArray(tc.input); got != tc.expected {
				t.Errorf("maxSubArray(%v) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
