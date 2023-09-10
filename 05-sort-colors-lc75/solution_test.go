package sortcolorslc75

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"Example 1",
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			"Example 2",
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sortColors(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("sortColors(%v) = %v; want = %v", tc.input, tc.input, tc.expected)
			}
		})
	}
}
