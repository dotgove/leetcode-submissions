package nextpermutationlc31

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Example 1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"Example 2", []int{3, 2, 1}, []int{1, 2, 3}},
		{"Example 3", []int{1, 1, 5}, []int{1, 5, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]int, len(tc.input))
			copy(original, tc.input)
			nextPermutation(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("nextPermutation(%v) = %v; wanted = %v", original, tc.input, tc.expected)
			}
		})
	}
}
