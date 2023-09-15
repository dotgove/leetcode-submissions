package pascalstrianglelc118

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected [][]int
	}{
		{
			"Example 1",
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			"Example 2",
			1,
			[][]int{{1}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := generate(tc.input); !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("generate(%d) = %v; want = %v", tc.input, got, tc.expected)
			}
		})
	}
}
