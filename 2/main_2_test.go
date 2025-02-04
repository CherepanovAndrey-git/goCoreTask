package main

import (
	"2/pkg"
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "mixed numbers with zero",
			input:  []int{1, 2, 3, 4, 0},
			expect: []int{2, 4},
		},
		{
			name:   "empty slice",
			input:  []int{},
			expect: []int{},
		},
		{
			name:   "only odd numbers",
			input:  []int{1, 3, 5},
			expect: []int{},
		},
		{
			name:   "even numbers with zero",
			input:  []int{2, 4, 6, 0},
			expect: []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pkg.SliceExample(tt.input)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("Expected %v, got %v", tt.expect, result)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		num    int
		expect []int
	}{
		{
			name:   "add to non-empty slice",
			input:  []int{1, 2},
			num:    3,
			expect: []int{1, 2, 3},
		},
		{
			name:   "add to empty slice",
			input:  []int{},
			num:    5,
			expect: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := pkg.AddElements(tt.input, tt.num)

			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("Expected %v, got %v", tt.expect, result)
			}

			if !reflect.DeepEqual(original, tt.input) {
				t.Error("Original slice was modified")
			}

			if len(result) != len(tt.input)+1 {
				t.Errorf("Expected length %d, got %d", len(tt.input)+1, len(result))
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "normal slice",
			input: []int{1, 2, 3},
		},
		{
			name:  "empty slice",
			input: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pkg.CopySlice(tt.input)

			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("Expected %v, got %v", tt.input, result)
			}

			if len(tt.input) > 0 {
				result[0] = 999
				if tt.input[0] == 999 {
					t.Error("Original slice was modified when changing copy")
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		index  int
		expect []int
	}{
		{
			name:   "remove middle element",
			input:  []int{1, 2, 3},
			index:  1,
			expect: []int{1, 3},
		},
		{
			name:   "remove first element",
			input:  []int{1, 2, 3},
			index:  0,
			expect: []int{2, 3},
		},
		{
			name:   "remove last element",
			input:  []int{1, 2, 3},
			index:  2,
			expect: []int{1, 2},
		},
		{
			name:   "single element slice",
			input:  []int{42},
			index:  0,
			expect: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := pkg.RemoveElement(tt.input, tt.index)

			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("Expected %v, got %v", tt.expect, result)
			}

			if !reflect.DeepEqual(original, tt.input) {
				t.Error("Original slice was modified")
			}

			if len(result) != len(tt.input)-1 {
				t.Errorf("Expected length %d, got %d", len(tt.input)-1, len(result))
			}
		})
	}
}
