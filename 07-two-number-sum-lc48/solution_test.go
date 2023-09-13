package twonumbersumlc48

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []int
		inputB   int
		expected []int
	}{
		{
			"Example 1",
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			"Example 2",
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			"Example 3",
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := twoSum(tc.inputA, tc.inputB); !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("twoSum(%v, %d) = %v; wanted = %v", tc.inputA, tc.inputB, got, tc.expected)
			}
		})
	}
}
