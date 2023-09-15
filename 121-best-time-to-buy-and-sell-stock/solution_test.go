package besttimetobuyandsellstocklc121

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Example 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"Example 2", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := maxProfit(tc.input); got != tc.expected {
				t.Errorf("maxProfit(%v) = %d; want = %d", tc.input, got, tc.expected)
			}
		})
	}
}
